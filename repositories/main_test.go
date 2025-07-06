package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB接続で使用する変数
var (
	testDB     *sql.DB
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

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
	// 環境変数の設定
	err := godotenv.Load("../docker-compose/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return err
	}
	if err := connectDB(); err != nil {
		fmt.Println("connect db", err)
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup", err)
		return err
	}
	return nil
}

// テスト後の後処理
func teardown() {
	cleanupDB()
	testDB.Close()
}

// DB接続処理
func connectDB() error {
	dbUser = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbDatabase)
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// DBセットアップ処理
func setupTestData() error {
	queryBytes, err := os.ReadFile("./testdata/setupDB.sql")
	if err != nil {
		return err
	}
	err = runQueries(queryBytes)
	return err
}

// DBクリーンアップ処理
func cleanupDB() error {
	queryBytes, err := os.ReadFile("./testdata/cleanupDB.sql")
	if err != nil {
		return err
	}
	err = runQueries(queryBytes)
	return err
}

// クエリ実行処理
func runQueries(queryBytes []byte) error {
	// 実行するクエリの準備
	queries := strings.Split(string(queryBytes), ";")
	for _, q := range queries {
		q = strings.TrimSpace(q)
		if q == "" {
			continue
		}
		_, err := testDB.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}
