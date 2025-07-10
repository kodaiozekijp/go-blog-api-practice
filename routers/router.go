package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodaiozekijp/go-blog-api-practice/controllers"
)

// ルータを作成し、返却する
func NewRouter(con *controllers.MyAppController) *mux.Router {
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

	// 作成したルータを返却
	return r
}
