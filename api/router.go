package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodaiozekijp/go-blog-api-practice/controllers"
	"github.com/kodaiozekijp/go-blog-api-practice/services"
)

// ルータを作成し、返却する
func NewRouter(db *sql.DB) *mux.Router {
	// MyAppService構造体を生成
	ser := services.NewMyAppService(db)

	// ArticleController及びCommentController構造体を生成
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

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
