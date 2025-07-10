package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	con := controllers.NewMyAppController(ser)

	// gorilla/muxのルータを使用
	r := mux.NewRouter()

	// ハンドラの登録
	r.HandleFunc("/", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article",
		con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list",
		con.ArticleListHandler).Methods(http.MethodGet)
	// 記事IDをパスパラメータのidとして取得
	r.HandleFunc("/article/{id:[0-9]+}",
		con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	// サーバ起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
