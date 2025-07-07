package services

import (
	"github.com/kodaiozekijp/go-blog-api-practice/models"
	"github.com/kodaiozekijp/go-blog-api-practice/repositories"
)

// 受け渡されたコメントを登録し、登録されたコメント情報を返却する
func PostCommentService(comment models.Comment) (models.Comment, error) {
	// sql.DB型の取得
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	// repositories層の関数InsertCommentでコメントを登録
	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
