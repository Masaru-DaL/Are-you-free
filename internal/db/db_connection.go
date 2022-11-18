package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/* Create MySQL Connection*/
func CreateConnection() *sql.DB {
	// docker-compose.ymlの環境変数の読み込みなど
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)

	db, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	// defer db.Close()
	/* 接続が可能であることを確認する */
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

// var Db *sql.DB

// /*main関数より早く呼び出し、MySQLの準備が完了し接続されるまで2秒ごとに接続を試みる*/
// func init() {
// 	// 環境変数を読み込み
// 	user := os.Getenv("MYSQL_USER")
// 	pw := os.Getenv("MYSQL_PASSWORD")
// 	db_name := os.Getenv("MYSQL_DATABASE")
// 	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
// 	var err error

// 	if Db, err = sql.Open("mysql", path); err != nil {
// 		log.Fatal("Db open error:", err.Error())
// 	}
// 	checkConnect(100)

// 	fmt.Println("db connected!!")
// }

// /* 2秒ごとに接続を行う再帰関数 */
// func checkConnect(count uint) {
// 	var err error
// 	if err = Db.Ping(); err != nil {
// 		time.Sleep(time.Second * 2)
// 		count--
// 		fmt.Printf("retry... count:%v\n", count)
// 		checkConnect(count)
// 	}
// }
