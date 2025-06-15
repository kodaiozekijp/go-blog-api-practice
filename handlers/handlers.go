package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	// URLからクエリパラメータを取得
	queryMap := req.URL.Query()
	// ページ番号に該当するデータを返す
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		// クエリパラメータがURLに含まれていない場合は、1ページ目のデータを返す
	} else {
		page = 1
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

// ArticleDetailHandler
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// URLからパスパラメータである記事IDを取得し、該当する記事を返す
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
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
