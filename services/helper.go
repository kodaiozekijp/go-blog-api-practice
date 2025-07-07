package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB接続で使用する変数
var (
	dbUser, dbPassword, dbDatabase = initEnv()
	dbConn                         = fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
)

// 環境変数から必要な値を取得する処理
func initEnv() (string, string, string) {
	// 環境変数を設定
	err := godotenv.Load("../docker-compose/.env")
	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	// 環境変数から必要な値を取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbDatabase := os.Getenv("MYSQL_DATABASE")

	return dbUser, dbPassword, dbDatabase
}

// DBに接続し、sql.DB型を取得する処理
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
