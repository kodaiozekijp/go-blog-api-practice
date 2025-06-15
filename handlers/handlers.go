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
	if req.Method == http.MethodGet {
		io.WriteString(w, "hello, World!\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// PostArticleHandler
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// ArticleListHandler
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// ArticleDetailHandler
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// PostNiceHandler
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// PostCommentHandler
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
