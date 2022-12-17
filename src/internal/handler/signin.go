package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"src/internal/auth"
	"src/internal/config"
	"src/internal/db"
	"src/internal/models"

	"github.com/labstack/echo/v4"
)

/* ユーザの新規作成 */
func SignUp(c echo.Context) error {
	db := db.CreateConnection()

	signUpName := c.FormValue("name")
	signUpPassword := c.FormValue("password")

	// ユーザ名をチェックし、存在した場合はエラー
	// 存在しなかったらユーザの新規作成を行う
	sqlStatement := "SELECT name, password FROM users WHERE name = ? AND password = ?"
	selectExecuteError := db.QueryRow(sqlStatement, signUpName, signUpPassword).Scan(&signUpName, &signUpPassword)
	switch {
	case selectExecuteError == sql.ErrNoRows:
		// ユーザを作成する
		sqlStatement := "INSERT INTO users(name, password) values(?, ?)"
		_, err := db.Exec(sqlStatement, signUpName, signUpPassword)

		if err != nil {
			log.Fatal(err)
		}
		// return c.JSON(http.StatusCreated, "created user name is:"+signUpName)
		return c.Redirect(http.StatusFound, "/login")
	case selectExecuteError != nil:
		log.Fatal(selectExecuteError)
		// log.Fatalf("query error: %v\n, err")
	}

	return c.JSON(http.StatusUnprocessableEntity, "The username you have entered already exists")
}

/* ユーザのサインイン機能 */
func SignIn(c echo.Context) error {
	db := db.CreateConnection()

	signInName := c.FormValue("name")
	signInPassword := c.FormValue("password")

	// ユーザ名、パスワードが正しいか確認する => トークンを作成して渡す
	sqlStatement := "SELECT name, password, id FROM users WHERE name = ? AND password = ?"

	rows, err := db.Query(sqlStatement, signInName, signInPassword)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	storedUser := new(models.User)
	if err := c.Bind(storedUser); err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&storedUser.Name, &storedUser.Password, &storedUser.ID)

		if err != nil {
			log.Fatal(err)
		}
	}

	if storedUser.ID == 0 && signInName == storedUser.Name {
		return c.JSON(http.StatusUnprocessableEntity, "The username and password you have entered is not correct.")
	} else {
		if storedUser.Password == signInPassword {
			// アクセストークンを作成する
			accessToken, exp, err := auth.GenerateAccessToken(storedUser)
			if err != nil {
				return err
			}
			// このクッキーはアクセスするための認証情報をもつ
			auth.SetTokenCookie(config.Config.AUTH.AccessTokenCookieName, accessToken, exp, c)
			// このクッキーはユーザの名前のみの情報を持つ
			auth.SetUserCookie(storedUser, exp, c)

			// refreshToken, exp, err := auth.GenerateRefreshToken(u)
			// if err != nil {
			// 	return err
			// }
			// auth.SetTokenCookie(config.Config.AUTH.RefreshTokenCookieName, refreshToken, exp, c)

			// return c.JSON(http.StatusOK, echo.Map{
			// 	"message":      "login successful",
			// 	"access_token": accessToken,
			// 	// "refresh_token": refreshToken,
			// })
			// return c.JSON(http.StatusOK, accessToken)
			return c.Redirect(http.StatusFound, "/auth/schedules")
		} else {
			return c.JSON(http.StatusUnprocessableEntity, "The username and password you have entered is not correct.")
		}
	}
}
