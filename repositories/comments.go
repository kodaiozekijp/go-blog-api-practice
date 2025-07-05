package repositories

import (
	"database/sql"

	"github.com/kodaiozekijp/go-blog-api-practice/models"
)

// 新規投稿をデータベースにINSERTする
// -> データベースに保存したコメント内容と、発生したエラーを戻り値にする
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	// INSERT文の定義
	const sqlStr = `
		INSERT INTO comments (article_id, message, created_at)
		VALUES (?, ?, now());
	`
	// 返却用のComment構造体の定義
	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	// INSERT文の実行
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}
	// 実行結果より、コメントIDの取得
	newCommentID, _ := result.LastInsertId()
	newComment.CommentID = int(newCommentID)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する
// -> 取得したコメント一覧と、発生したエラーを戻り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	// SELECT文の定義
	const sqlStr = `
		SELECT * FROM comments WHERE article_id = ?;
	`
	// SELECT文の実行
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 返却用のmodels.Comment構造体のスライスを定義
	commentArray := make([]models.Comment, 0)
	// 取得したレコードをcommentArrayに格納する
	for rows.Next() {
		var comment models.Comment
		var createdAt sql.NullTime
		err = rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message,
			&createdAt)
		if err != nil {
			return nil, err
		}
		if createdAt.Valid {
			comment.CreatedAt = createdAt.Time
		}
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
