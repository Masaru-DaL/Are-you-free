package handler

import (
	"fmt"
	"net/http"
	"src/internal/db"
	"src/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

	user := models.User{}
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

	users := models.Users{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Is_admin, &user.Created_at, &user.Updated_at)

		if err != nil {
			fmt.Println(err)
		}
		users.Users = append(users.Users, user)
	}
	return c.JSON(http.StatusOK, users)
}

/* userの名前、パスワードを更新する(自身の持つIDが分かる必要がある) */
func UpdateUser(c echo.Context) error {
	db := db.CreateConnection()

	user := new(models.User)
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

	user := new(models.User)
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
