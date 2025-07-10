package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodaiozekijp/go-blog-api-practice/controllers"
)

// ルータを作成し、返却する
func NewRouter(aCon *controllers.ArticleController, cCon *controllers.CommentController) *mux.Router {
	// gorilla/muxのルータを使用
	r := mux.NewRouter()

	// ハンドラの登録
	r.HandleFunc("/", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article",
		aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list",
		aCon.ArticleListHandler).Methods(http.MethodGet)
	// 記事IDをパスパラメータのidとして取得
	r.HandleFunc("/article/{id:[0-9]+}",
		aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	// 作成したルータを返却
	return r
}
