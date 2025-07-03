package repositories_test

import (
	"testing"

	"github.com/kodaiozekijp/go-blog-api-practice/models"
	"github.com/kodaiozekijp/go-blog-api-practice/repositories"
)

func TestSelectArticleList(t *testing.T) {
	// テスト対象の関数の実行
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	// SelectArticleList関数から得たArticleスライスの長さと期待値の検証
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
