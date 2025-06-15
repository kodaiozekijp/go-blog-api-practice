package main

import (
	"log"
	"net/http"

	"github.com/kodaiozekijp/go-blog-api-practice/handlers"
)

func main() {
	// ハンドラの登録
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバ起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
