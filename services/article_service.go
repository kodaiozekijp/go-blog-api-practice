package services

import (
	"github.com/kodaiozekijp/go-blog-api-practice/models"
	"github.com/kodaiozekijp/go-blog-api-practice/repositories"
)

// 受け渡された記事を登録し、登録された記事情報を返却する
func PostArticleService(article models.Article) (models.Article, error) {
	// sql.DB型の取得
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// repositories層の関数InsertArticleで記事を登録
	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

// 指定されたページの記事の一覧を返却する
func GetArticleListService(page int) ([]models.Article, error) {
	// sql.DB型の取得
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// repositories層の関数SelectArticleListで指定されたページの記事一覧を取得
	articles, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// 指定されたIDの記事を返却
func GetArticleService(articleID int) (models.Article, error) {
	// sql.DB型の取得
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// repositories層の関数SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 取得したコメント一覧を、取得した記事に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// 指定された記事のいいね数を1増やして、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	// sql.DB型の取得
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// repositories層の関数UpdateNiceNumでいいね数を1増やす
	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		Author:    article.Author,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
