package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kodaiozekijp/go-blog-api-practice/apperrors"
	"github.com/kodaiozekijp/go-blog-api-practice/controllers/services"
	"github.com/kodaiozekijp/go-blog-api-practice/models"
)

// コントローラ構造体を定義
type ArticleController struct {
	// フィールドにArticleServicerインタフェースを持たせる
	service services.ArticleServicer
}

// コンストラクタの定義
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// ハンドラ定義
// HelloHandler
func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	// ハンドラの処理内容
	io.WriteString(w, "hello, World!\n")
}

// PostArticleHandler
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// jsonのデコーダーを使用し、リクエストボディをデコードし、記事を取得する
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	// services層の関数PostArticleServiceで記事を登録する
	newArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// jsonのエンコーダーを使用し、記事をエンコードした上で、返却する
	json.NewEncoder(w).Encode(newArticle)
}

// ArticleListHandler
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// URLからクエリパラメータを取得
	queryMap := req.URL.Query()
	// ページ番号に該当するデータを返す
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParameter.Wrap(err, "query parameter must be number")
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		// クエリパラメータがURLに含まれていない場合は、1ページ目のデータを返す
	} else {
		page = 1
	}

	// services層の関数GetArticleListServiceで記事の一覧を取得
	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// jsonのエンコーダーを使用し、記事一覧をjsonにエンコードした上で返却する
	json.NewEncoder(w).Encode(articleList)
}

// ArticleDetailHandler
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// URLからパスパラメータである記事IDを取得し、該当する記事を返す
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParameter.Wrap(err, "path parameter must be number")
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	//  services層の関数GetArticleServiceでIDに紐づく記事を取得
	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// jsonのエンコーダーを使用し、記事1をjsonにエンコードした上で、返却する
	json.NewEncoder(w).Encode(article)
}

// PostNiceHandler
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// jsonのデコーダーを使用し、リクエストボディをデコードし、記事を取得する
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode\n", http.StatusBadRequest)
		return
	}

	// services層の関数PostNiceServiceで記事のいいね数を1増やす
	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// jsonのエンコーダーを使用し、記事をjsonにエンコードした上で、返却する
	json.NewEncoder(w).Encode(article)
}
