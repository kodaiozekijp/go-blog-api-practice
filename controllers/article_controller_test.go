package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArticleListHandler(t *testing.T) {
	// テスト対象のハンドラメソッドに入れるinputを定義
	var tests = []struct {
		title      string
		query      string
		resultCode int
	}{
		{title: "number query test", query: "1", resultCode: http.StatusOK},
		{title: "alphabet query test", query: "aaa",
			resultCode: http.StatusBadRequest},
	}

	// テストの実行
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// ハンドラに渡す引数の用意
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s",
				test.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			// ハンドラメソッドに用意した引数を受け渡す
			aCon.ArticleListHandler(res, req)

			// outputと期待値との検証
			if res.Code != test.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", test.resultCode, res.Code)
			}
		})
	}
}
