package model

import (
	"fmt"
	"os"
)

/* DBのdsnを取得する関数 */
func GetDBConfig() string {
	// docker-compose.ymlの環境変数を読み込む
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DBNAME")

	// dsn: (DBの接続情報に付ける識別子)の定義
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}
