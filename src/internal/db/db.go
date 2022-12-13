package db

import (
	"database/sql"
	"fmt"
	"src/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

/* Create MySQL Connection*/
func CreateConnection() *sql.DB {
	db, err := sql.Open(config.Config.DB.SQLDriver, config.Config.DB.Path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("MySQL is connected")
	}
	// defer db.Close()

	/* 接続が可能であることを確認する */
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL is not connected")
		fmt.Println(err.Error())
	}
	return db
}
