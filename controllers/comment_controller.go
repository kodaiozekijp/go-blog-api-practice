package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kodaiozekijp/go-blog-api-practice/controllers/services"
	"github.com/kodaiozekijp/go-blog-api-practice/models"
)

// コントローラ構造体を定義
type CommentController struct {
	// フィールドにCommentServicerインタフェースを持たせる
	service services.CommentServicer
}

// コンストラクタの定義
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// ハンドラ定義
// PostCommentHandler
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonのデコーダーを使用し、リクエストボディをデコードし、コメントを取得する
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode\n", http.StatusBadRequest)
		return
	}

	//
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// jsonのエンコーダーを使用し、コメントをjsonにエンコードした上で、返却する
	json.NewEncoder(w).Encode(comment)
}
