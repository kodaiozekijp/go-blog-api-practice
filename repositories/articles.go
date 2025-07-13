package repositories

import (
	"database/sql"

	"github.com/kodaiozekijp/go-blog-api-practice/models"
)

const (
	articleNumPerPage = 5
)

// 新規投稿をデータベースにINSERTする
// -> データベースに保存した記事内容と、発生したエラーを戻り値にする
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	// INSERT文の定義
	const sqlStr = `
		INSERT INTO articles (title, contents, author, nice, created_at)
		VALUES (?, ?, ?, 0, now());
	`
	// 返却用記事の定義
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.Author =
		article.Title, article.Contents, article.Author
	// INSERT文の実行
	result, err := db.Exec(sqlStr, article.Title, article.Contents,
		article.Author)
	if err != nil {
		return models.Article{}, err
	}
	// 記事IDの取得
	newArticleID, _ := result.LastInsertId()
	newArticle.ID = int(newArticleID)

	return newArticle, nil
}

// 変数pageで指定されたページに表示する投稿一覧をデータベースからSELECTする
// -> 取得した記事と、発生したエラーを戻り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	// SELECT文の定義
	const sqlStr = `
		SELECT article_id, title, contents, author, nice FROM articles
		LIMIT ? OFFSET ?;
	`
	// SELECT文の実行
	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 返却用のmodels.Article構造体のスライスの定義
	articleList := make([]models.Article, 0)
	// 取得したレコードをarticleArrayに格納する
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Contents,
			&article.Author, &article.NiceNum)
		if err != nil {
			return nil, err
		}
		articleList = append(articleList, article)
	}

	return articleList, nil
}

// 記事IDを指定して、記事を取得する
// -> 取得した記事と、発生したエラーを戻り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	// SELECT文の定義
	const sqlStr = `
		SELECT * FROM articles WHERE article_id = ?;
	`
	// SELECT文の実行
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {

		return models.Article{}, err
	}
	// 返却用の記事の定義
	var article models.Article
	// 取得したレコードをarticleに格納（作成日時はNULLの可能性を考慮）
	var createdAt sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents,
		&article.Author, &article.NiceNum, &createdAt)
	if err != nil {
		return models.Article{}, err
	}
	// 作成日時がNULLでない場合は、articleに格納
	if createdAt.Valid {
		article.CreatedAt = createdAt.Time
	}

	return article, nil
}

// いいねの数を+1する
// -> 発生したエラーを戻り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	// トランザクションを開始する
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// いいね数を取得するSELECT文の定義
	const sqlGetNice = `
		SELECT nice FROM articles WHERE article_id = ?;
	`
	// SELECT文の実行
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	// 取得した値を変数niceNumに格納
	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	// いいね数を+1するUPDATE文の定義
	const sqlUpdateNice = `
		UPDATE articles SET nice = ? WHERE article_id = ?;
	`
	// UPDATE文の実行
	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}
	// トランザクションをコミットする
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
