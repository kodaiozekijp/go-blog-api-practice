package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodaiozekijp/go-blog-api-practice/handlers"
)

func main() {
	// gorilla/muxのルータを使用
	r := mux.NewRouter()

	// ハンドラの登録
	r.HandleFunc("/", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// サーバ起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
