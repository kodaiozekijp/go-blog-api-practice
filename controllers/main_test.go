package controllers_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kodaiozekijp/go-blog-api-practice/controllers"
	"github.com/kodaiozekijp/go-blog-api-practice/services"
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
	err := godotenv.Load("./docker-compose/.env")
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

// テストに使うコントローラ構造体を定義
var aCon *controllers.ArticleController
var cCon *controllers.CommentController

func TestMain(m *testing.M) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("DB setup fail")
		os.Exit(1)
	}

	ser := services.NewMyAppService(db)
	aCon = controllers.NewArticleController(ser)
	cCon = controllers.NewCommentController(ser)

	m.Run()
}
