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

	// テスト結果として期待する値を定義
	expected := models.Article{
		ID:       1,
		Title:    "1st Post",
		Contents: "This is my first blog",
		Author:   "kodai",
		NiceNum:  1,
	}

	// DBからarticle_id=1の記事を取得する
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	// 取得した値と期待値との検証
	if got.ID != expected.ID {
		t.Errorf("ID: got %d but want %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("Title: got %s but want %s\n", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("Contents: got %s but want %s\n", got.Contents, expected.Contents)
	}
	if got.Author != expected.Author {
		t.Errorf("Author: got %s but want %s\n", got.Author, expected.Author)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}
