package models

import (
	"fmt"
	"net/http"
	"src/internal/db"
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

/* POSTリクエスト */
func CreateUser(c echo.Context) error {
	con := db.CreateConnection()
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	sqlStatement := "INSERT INTO users(name, password) VALUES(?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(user.Name, user.Password)

	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	return c.JSON(http.StatusCreated, user.)
}
