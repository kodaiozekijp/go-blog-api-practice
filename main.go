package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kodaiozekijp/go-blog-api-practice/controllers"
	"github.com/kodaiozekijp/go-blog-api-practice/routers"
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

func main() {
	// DBへの接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	// MyAppService構造体を生成
	ser := services.NewMyAppService(db)

	// MyAppController構造体を生成
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	// ルータを生成
	rou := routers.NewRouter(aCon, cCon)

	// サーバ起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", rou))
}
