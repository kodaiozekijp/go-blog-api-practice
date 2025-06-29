package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kodaiozekijp/go-blog-api-practice/handlers"
	"github.com/kodaiozekijp/go-blog-api-practice/models"
)

func main() {
	// gorilla/muxのルータを使用
	r := mux.NewRouter()

	// ハンドラの登録
	r.HandleFunc("/", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article",
		handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list",
		handlers.ArticleListHandler).Methods(http.MethodGet)
	// 記事IDをパスパラメータのidとして取得
	r.HandleFunc("/article/{id:[0-9]+}",
		handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// サーバ起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", r))

	// DB(mysql)への接続
	// 接続情報の宣言
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "blog_api_db"
	// データベースに接続するためのアドレス文を定義
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbDatabase)

	// Open関数を用いてデータベースに接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	// プログラムが終了するときに、コネクションがcloseされるようにする
	defer db.Close()

	// レコードの取得
	// クエリの準備
	const sqlStr = `
			SELECT * FROM articles WHERE article_id = ?;
	`
	// クエリを実行し、レコードを取得する
	articleID := 1
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// rowのレコードを構造体に格納する
	// レコードの各カラムを構造体に格納する
	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(&article.ID, &article.Title, &article.Contents,
		&article.Author, &article.NiceNum, &createdTime)
	// エラーの場合は、retrunする
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得したレコードのcreated_atがnullでない場合は構造体に格納
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	fmt.Printf("%+v\n", article)

	// レコードの挿入
	insertArticle := models.Article{
		Title:    "insert test",
		Contents: "can I insert data correctly?",
		Author:   "kodai",
	}
	const sqlStrIns = `
		INSERT INTO articles (title, contents, author, nice, created_at)
		VALUES (?, ?, ?, 0, now());
	`
	// INSERT文を実行する
	result, err := db.Exec(sqlStrIns, insertArticle.Title, insertArticle.Contents,
		insertArticle.Author)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 結果を確認
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// トランザクションを利用したレコードの更新
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 更新する記事を取得する
	updateArticleID := 2
	const sqlStrGetNice = `
		SELECT nice FROM articles WHERE article_id = ?;
	`
	// SELECT文を実行する
	updateRow := tx.QueryRow(sqlStrGetNice, updateArticleID)
	if err := updateRow.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	// 変数niceNumに現在のいいね数を保持する
	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	// いいね数を+1する更新処理
	const sqlStrUpdateNice = `
		UPDATE articles SET nice = ? WHERE article_id = ?;
	`
	// UPDATE文を実行する
	_, err = tx.Exec(sqlStrUpdateNice, updateArticleID, niceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	// コミットして処理内容を確定させる
	tx.Commit()
}
