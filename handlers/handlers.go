package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// ハンドラ定義
// HelloHandler
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// ハンドラの処理内容
	io.WriteString(w, "hello, World!\n")
}

// PostArticleHandler
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// ArticleListHandler
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

// ArticleDetailHandler
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// PostNiceHandler
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// PostCommentHandler
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
