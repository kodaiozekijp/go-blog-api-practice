package repositories_test

import (
	"testing"

	"github.com/kodaiozekijp/go-blog-api-practice/models"
	"github.com/kodaiozekijp/go-blog-api-practice/repositories"
)

func TestInsertComment(t *testing.T) {
	// コメントの定義
	comment := models.Comment{
		ArticleID: 2,
		Message:   "1st comment to 2nd Post",
	}

	// テスト対象の関数の実行
	var expectedCommentID = 3
	got, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	// 取得した結果と期待値との検証
	if got.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID,
			got.CommentID)
	}

	// テスト後処理（追加したコメントの削除）
	t.Cleanup(func() {
		// 削除クエリの実行
		const sqlDeleteComment = `
			DELETE FROM comments WHERE article_id = ? AND message = ?;
		`
		testDB.Exec(sqlDeleteComment, comment.ArticleID, comment.Message)
	})
}

func TestSelectCommentList(t *testing.T) {
	// テスト対象の関数の実行
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	// 取得したコメントの個数と期待値との検証
	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID,
				comment.ArticleID)
		}
	}
}
