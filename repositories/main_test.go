package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// テスト全体で共有するsql.DB型
var testDB *sql.DB

// テストを実行する
func TestMain(m *testing.M) {
	// 前処理の呼び出し
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	// テスト実行
	m.Run()

	// 後処理の呼び出し
	teardown()
}

// 全テスト共通の前処理
func setup() error {
	// DBへの接続
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "blog_api_db"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// テスト後の後処理
func teardown() {
	testDB.Close()
}
