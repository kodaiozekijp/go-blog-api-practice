package services

import (
	"database/sql"
	"errors"

	"github.com/kodaiozekijp/go-blog-api-practice/apperrors"
	"github.com/kodaiozekijp/go-blog-api-practice/models"
	"github.com/kodaiozekijp/go-blog-api-practice/repositories"
)

// サービス構造体の定義
type MyAppService struct {
	// フィールドにsql.DB型を持たせる
	db *sql.DB
}

// コンストラクタの定義
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}

// PostArticleHandlerで使うことを想定したサービス
// 受け渡された記事を登録し、登録された記事情報を返却する
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	// repositories層の関数InsertArticleで記事を登録
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		// 独自エラーのMyAppErrorでerrorをラップする
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定されたページの記事の一覧を返却する
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	// repositories層の関数SelectArticleListで指定されたページの記事一覧を取得
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		// 独自エラーのMyAppErrorでerrorをラップする
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

// ArticleDetailHandlerで使うことを想定したサービス
// 指定されたIDの記事を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		// 取得したデータが0件か確認
		if errors.Is(err, sql.ErrNoRows) {
			// 独自エラーのMyAppErrorでerrorをラップする
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		// 独自エラーのMyAppErrorでerrorをラップする
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	// repositories層の関数SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		// 独自エラーのMyAppErrorでerrorをラップする
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	// 取得したコメント一覧を、取得した記事に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostNiceHandlerで使うことを想定したサービス
// 指定された記事のいいね数を1増やして、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	// repositories層の関数UpdateNiceNumでいいね数を1増やす
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		// 更新対象のデータが存在するか確認
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		// 独自エラーのMyAppErrorでerrorをラップする
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice num")
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

// PostCommentHandlerで使うことを想定したサービス
// 受け渡されたコメントを登録し、登録されたコメント情報を返却する
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	// repositories層の関数InsertCommentでコメントを登録
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		// 独自エラーのMyAppErrorでerrorをラップする
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
