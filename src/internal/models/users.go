package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"src/internal/auth"
	"src/internal/db"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Is_admin   bool      `json:"is_admin"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Users struct {
	Users []User `json:"users"`
}

/* ユーザのサインイン機能 */
func SignIn(c echo.Context) error {
	db := db.CreateConnection()

	tmpUser := new(User)
	if err := c.Bind(tmpUser); err != nil {
		return err
	}

	// ユーザ名、パスワードが正しいか確認する => トークンを作成して渡す
	sqlStatement := "SELECT name, password, id FROM users WHERE name = ? AND password = ?"

	rows, err := db.Query(sqlStatement, tmpUser.Name, tmpUser.Password)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var username string
	var password string
	var id int
	for rows.Next() {
		err := rows.Scan(&username, &password, &id)

		if err != nil {
			log.Fatal(err)
		}
	}

	if username == "" || password == "" {
		return c.JSON(http.StatusUnprocessableEntity, "The username and password you have entered is not correct.")
	} else {
		// パスワードをチェックする
		if password == tmpUser.Password {
			token := auth.GenerateToken(id)

			http.SetCookie(c.Response().Writer, &http.Cookie{
				Name:  "token",
				Value: token,
			})
			return c.JSON(http.StatusOK, token)
		} else {
			return c.JSON(http.StatusUnprocessableEntity, "The username and password you have entered is not correct.")
		}
	}
}

/* user情報を1件取得する */
func GetUser(c echo.Context) error {
	db := db.CreateConnection()

	user_id := c.Param("id")
	strconv.Atoi(user_id)

	sqlStatement := "SELECT * FROM users WHERE id = ? LIMIT 1"

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	user := User{}
	err2 := stmt.QueryRow(user_id).Scan(&user.ID, &user.Name, &user.Password, &user.Is_admin, &user.Created_at, &user.Updated_at)
	if err2 != nil {
		fmt.Println(err2)
	}

	return c.JSON(http.StatusOK, user)
}

/* user情報を全件取得する */
func GetUsers(c echo.Context) error {
	db := db.CreateConnection()

	sqlStatement := "SELECT * FROM users"

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	// rows, err := stmt.Query(1)
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	users := Users{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Is_admin, &user.Created_at, &user.Updated_at)

		if err != nil {
			fmt.Println(err)
		}
		users.Users = append(users.Users, user)
	}
	return c.JSON(http.StatusOK, users.Users)
}

/* ユーザの新規作成 */
func SignUp(c echo.Context) error {
	db := db.CreateConnection()

	tmpUser := new(User)
	if err := c.Bind(tmpUser); err != nil {
		return err
	}

	// ユーザ名をチェックし、存在した場合はエラー
	// 存在しなかったらユーザの新規作成を行う
	sqlStatement := "SELECT name, password FROM users WHERE name = ? AND password = ?"
	selectExecuteError := db.QueryRow(sqlStatement, tmpUser.Name, tmpUser.Password).Scan(&tmpUser.Name, &tmpUser.Password)
	switch {
	case selectExecuteError == sql.ErrNoRows:
		// ユーザを作成する
		sqlStatement := "INSERT INTO users(name, password) values(?, ?)"
		_, err := db.Exec(sqlStatement, tmpUser.Name, tmpUser.Password)

		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusCreated, tmpUser)
	case selectExecuteError != nil:
		log.Fatal(selectExecuteError)
		// log.Fatalf("query error: %v\n, err")
	}

	return c.JSON(http.StatusUnprocessableEntity, "The username you have entered already exists")
}

/* userの名前、パスワードを更新する(自身の持つIDが分かる必要がある) */
func UpdateUser(c echo.Context) error {
	db := db.CreateConnection()

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	sqlStatement := "UPDATE users SET name=?, password=? WHERE id=?"
	stmt, err := db.Prepare(sqlStatement)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Password, user.ID)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		return c.JSON(http.StatusCreated, user)
	}

	return c.JSON(http.StatusOK, user.ID)
}

/* Admin権限でIDを指定してuserの情報の全てを更新する */
func UpdateUserByAdmin(c echo.Context) error {
	db := db.CreateConnection()

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	sqlStatement := "UPDATE users SET name=?, password=?, is_admin=?, deposit=? WHERE id=?"
	stmt, err := db.Prepare(sqlStatement)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Password, user.Is_admin, user.ID)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		return c.JSON(http.StatusCreated, user)
	}

	return c.JSON(http.StatusOK, user.ID)
}

/* Admin権限でIDを指定してuserの情報を削除する */
func DeleteUserByAdmin(c echo.Context) error {
	db := db.CreateConnection()

	request_id := c.Param("id")
	sqlStatement := "DELETE FROM users where id = ?"
	stmt, err := db.Prepare(sqlStatement)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(request_id)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.RowsAffected())
	return c.JSON(http.StatusOK, "Deleted")
}
