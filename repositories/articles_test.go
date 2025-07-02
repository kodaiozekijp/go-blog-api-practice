package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/kodaiozekijp/go-blog-api-practice/models"
	"github.com/kodaiozekijp/go-blog-api-practice/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	// DBへの接続
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "blog_api_db"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

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
			got, err := repositories.SelectArticleDetail(db, test.expected.ID)
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
