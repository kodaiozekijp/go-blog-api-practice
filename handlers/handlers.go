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
	// jsonのデコーダーを使用し、リクエストボディをデコードし、記事を取得する
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	article := reqArticle

	// jsonのエンコーダーを使用し、記事をエンコードした上で、返却する
	json.NewEncoder(w).Encode(article)
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

	// jsonのエンコーダーを使用し、記事1と記事2をjsonにエンコードした上で、返却する
	articleList := []models.Article{models.Article1, models.Article2}

	json.NewEncoder(w).Encode(articleList)
}

// ArticleDetailHandler
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// URLからパスパラメータである記事IDを取得し、該当する記事を返す
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// articleIDについては後程使用する為、暫定処理を追加
	log.Println(articleID)

	// jsonのエンコーダーを使用し、記事1をjsonにエンコードした上で、返却する
	article := models.Article1

	json.NewEncoder(w).Encode(article)
}

// PostNiceHandler
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// jsonのデコーダーを使用し、リクエストボディをデコードし、記事を取得する
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode\n", http.StatusBadRequest)
		return
	}
	article := reqArticle

	// jsonのエンコーダーを使用し、記事をjsonにエンコードした上で、返却する
	json.NewEncoder(w).Encode(article)
}

// PostCommentHandler
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonのデコーダーを使用し、リクエストボディをデコードし、コメントを取得する
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode\n", http.StatusBadRequest)
		return
	}
	comment := reqComment

	// jsonのエンコーダーを使用し、コメントをjsonにエンコードした上で、返却する
	json.NewEncoder(w).Encode(comment)
}
