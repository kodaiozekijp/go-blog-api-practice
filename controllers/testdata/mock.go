package testdata

import (
	"github.com/kodaiozekijp/go-blog-api-practice/models"
)

// サービス構造体の定義
type serviceMock struct{}

// コンストラクタの定義
func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

// Article構造体を返却する
func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

// Article構造体を返却する
func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return articleTestData, nil
}

// Article構造体を返却する
func (s *serviceMock) GetArticleService(articleID int) (models.Article, error) {
	return articleTestData[0], nil
}

// Article構造体を返却する
func (s *serviceMock) PostNiceService(article models.Article) (models.Article, error) {
	return articleTestData[0], nil
}

// Comment構造体を返却する
func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}
