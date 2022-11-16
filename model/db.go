package model

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

/* DB接続とテーブルを作成する関数 */
func DBConnection() *sql.DB {
	// 関数GetDBConfigを実行し、戻り値をdsnと定義する
	dsn := GetDBConfig()
	// エラー型のerrを定義する
	var err error
	// dsnを使ってDBに接続する。戻り値をdbとerrに代入する
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	// テーブルの作成
	CreateTable(db)
	// *gorm.DB型を *sql.DB型に変換する
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	// sqlDBを返す
	return sqlDB
}

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
