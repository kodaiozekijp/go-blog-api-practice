package repositories_test

import (
	"testing"

	"github.com/kodaiozekijp/go-blog-api-practice/models"
	"github.com/kodaiozekijp/go-blog-api-practice/repositories"
)

func TestInsertArticle(t *testing.T) {
	// テスト対象の関数に渡す記事の定義
	article := models.Article{
		Title:    "insertTest",
		Contents: "insert test",
		Author:   "kodai",
		NiceNum:  0,
	}

	// テスト対象の関数の実行
	expectedID := 3
	got, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Fatal(err)
	}

	// InsertArticle関数から得た記事IDと期待値の検証
	if got.ID != expectedID {
		t.Errorf("new article id is expected %d but got %d\n", expectedID, got.ID)
	}

	// 個別の後処理の実装(追加したレコードの削除)
	t.Cleanup(func() {
		const sqlStr = `
			DELETE FROM articles WHERE title = ? AND contents = ? AND author = ?;
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.Author)
	})
}

func TestSelectArticleList(t *testing.T) {
	// テスト対象の関数の実行
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	// SelectArticleList関数から得た記事スライスの長さと期待値の検証
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	// テーブルドリブンテストの実装
	// struct構造体のスライスを準備
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "1st Post",
				Contents: "This is my first blog",
				Author:   "kodai",
				NiceNum:  1,
			},
		}, {
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd Post",
				Contents: "Second blog post",
				Author:   "kodai",
				NiceNum:  2,
			},
		},
	}

	// サブテストを実行する
	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			// DBから記事を取得する
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			// 取得した値と期待値との検証
			if got.ID != test.expected.ID {
				t.Errorf("ID: got %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: got %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: got %s but want %s\n", got.Contents,
					test.expected.Contents)
			}
			if got.Author != test.expected.Author {
				t.Errorf("Author: got %s but want %s\n", got.Author,
					test.expected.Author)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum,
					test.expected.NiceNum)
			}
		})
	}
}

func TestUpdateNiceNum(t *testing.T) {
	// 期待値の記事の定義
	expected := models.Article{
		ID:      1,
		NiceNum: 2,
	}
	// テスト対象の関数を実行
	err := repositories.UpdateNiceNum(testDB, expected.ID)
	if err != nil {
		t.Fatal(err)
	}
	// 関数実行後のいいね数を取得
	const sqlGetNice = `
		SELECT nice FROM articles WHERE article_id = ?;
	`
	var niceNum int
	row := testDB.QueryRow(sqlGetNice, expected.ID)
	if err = row.Err(); err != nil {
		t.Fatal(err)
	}
	// 取得したいいね数と期待値の比較検証
	if err = row.Scan(&niceNum); err != nil {
		t.Fatal(err)
	}
	if niceNum != expected.NiceNum {
		t.Errorf("nice num is expected %d but %d", expected.NiceNum, niceNum)
	}
	// 後処理
	t.Cleanup(func() {
		// いいね数を1減らすクエリの定義
		const sqlMinusNiceNum = `
			UPDATE articles SET nice = ? WHERE article_id = ?;
		`
		// クエリの実行
		testDB.Exec(sqlMinusNiceNum, niceNum-1, expected.ID)
	})
}
