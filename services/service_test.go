package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kodaiozekijp/go-blog-api-practice/services"
)

// DB接続で使用する変数
var (
	dbUser, dbPassword, dbDatabase = initEnv()
	dbConn                         = fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
)

// レシーバーとして使うためのMyAppService構造体
var ser *services.MyAppService

// 環境変数から必要な値を取得する処理
func initEnv() (string, string, string) {
	// 環境変数を設定
	err := godotenv.Load("../docker-compose/.env")
	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	// 環境変数から必要な値を取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbDatabase := os.Getenv("MYSQL_DATABASE")

	return dbUser, dbPassword, dbDatabase
}

func TestMain(m *testing.M) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sql.DB型を引数にMyAppService構造体を作成
	ser = services.NewMyAppService(db)

	// 個別のベンチマークテストの実行
	m.Run()
}

// GetArticleServiceのベンチマークテスト
func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	// 計測を始める為にタイマーを初期化
	b.ResetTimer()

	// GetArticleListServicesの実行
	for i := 0; i < b.N; i++ {
		_, err := ser.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
