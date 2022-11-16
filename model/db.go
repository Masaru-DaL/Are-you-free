package model

import (
	"fmt"
	"os"

	"gorm.io/gorm"
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

/* テーブルを作成する関数 */
func CreateTable(db *gorm.DB) {
	// 空き予定の開始テーブル
	db.AutoMigrate(&StartTime{})
	// 空き予定の終了テーブル
	db.AutoMigrate(&EndTime{})
}
