package controllers_test

import (
	"testing"

	"github.com/kodaiozekijp/go-blog-api-practice/controllers"
	"github.com/kodaiozekijp/go-blog-api-practice/controllers/testdata"
)

// テストに使うコントローラ構造体を定義
var aCon *controllers.ArticleController
var cCon *controllers.CommentController

func TestMain(m *testing.M) {
	// サービスのモックからコントローラを生成する
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	cCon = controllers.NewCommentController(ser)

	m.Run()
}
