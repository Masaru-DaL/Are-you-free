package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

/* DB接続とテーブルの作成を行う関数 */
func DBConnection() *sql.DB {
	dsn := GetDBConfig()
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	CreateTable(db)
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	return sqlDB
}

/* GetDBconfig: DBのdsnを取得する関数 */
func GetDBConfig() string {
	// docker-compose.ymlの環境変数を読み込む
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DBNAME")

	// dsnの定義
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=Local"

	return dsn
}

/* CreateTable: テーブルを作成する関数 */
func CreateTable(db *gorm.DB) {
	// // 空き予定の開始
	// db.AutoMigrate(&StartTime{})
	// // 空き予定の終了
	// db.AutoMigrate(&EndTime{})
	db.AutoMigrate(&Time{})
}
