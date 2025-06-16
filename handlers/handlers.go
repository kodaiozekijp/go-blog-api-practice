package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kodaiozekijp/go-blog-api-practice/models"
)

// ハンドラ定義
// HelloHandler
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// ハンドラの処理内容
	io.WriteString(w, "hello, World!\n")
}

// PostArticleHandler
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// 記事1をjsonにエンコードした上で、返却する
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
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

	// pageについては後程使用する為、暫定処理を追加
	log.Println(page)

	// 記事1と記事2をjsonにエンコードした上で、返却する
	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// ArticleDetailHandler
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// URLからパスパラメータである記事IDを取得し、該当する記事を返す
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 記事1をjsonにエンコードした上で、返却する
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}

	// articleIDについては後程使用する為、暫定処理を追加
	log.Println(articleID)

	w.Write(jsonData)
}

// PostNiceHandler
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// 記事1をjsonにエンコードした上で、返却する
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// PostCommentHandler
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// コメント1をjsonにエンコードした上で、返却する
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
