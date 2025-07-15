package controllers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestPostArticleHandler(t *testing.T) {
	// テストケースを用意
	var tests = []struct {
		title      string
		reqBody    string
		resultCode int
	}{
		{title: "decoded test", reqBody: `{"title": "aaa","contents": "bbb","author": "me","nice_num": 0,"comments": [],"created_at": "2024-01-01T00:00:00Z"}`, resultCode: http.StatusOK},
		{title: "can't be decoded test", reqBody: `{"title":,"contents":"ccc"}`, resultCode: http.StatusBadRequest},
	}

	// テストの実行
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// ハンドラに渡す引数の用意
			url := "http://localhost:8080/article"
			jsonBytes := []byte(test.reqBody)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
			req.Header.Set("Content-type", "application/json")

			res := httptest.NewRecorder()

			// ハンドラにメソッドに用意した引数を渡す
			aCon.PostArticleHandler(res, req)

			// outputと期待値との検証
			if res.Code != test.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", test.resultCode, res.Code)
			}
		})
	}
}

func TestArticleListHandler(t *testing.T) {
	// テストケースを用意
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

func TestArticleDetailHandler(t *testing.T) {
	// テストケースを用意
	var tests = []struct {
		title      string
		articleID  string
		resultCode int
	}{
		{title: "number pathparameter", articleID: "1", resultCode: http.StatusOK},
		{title: "alphabet pathparameter", articleID: "aaa", resultCode: http.StatusNotFound},
	}

	// テスト実行
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// ハンドラに渡す引数の用意
			url := fmt.Sprintf("http://localhost:8080/article/%s", test.articleID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			// ルータを作成し、ハンドラメソッドを呼び出す
			r := mux.NewRouter()
			r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			// outputと期待値との検証
			if res.Code != test.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", test.resultCode, res.Code)
			}
		})
	}
}

func TestPostNiceHandler(t *testing.T) {
	// テストケースを用意
	var tests = []struct {
		title      string
		reqBody    string
		resultCode int
	}{
		{title: "decoded test", reqBody: `{"title": "aaa","contents": "bbb","author": "me","nice_num": 0,"comments": [],"created_at": "2024-01-01T00:00:00Z"}`, resultCode: http.StatusOK},
		{title: "can't be decoded test", reqBody: `{"title":,"contents":"ccc"}`, resultCode: http.StatusBadRequest},
	}

	// テストの実行
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// ハンドラに渡す引数の用意
			url := "http://localhost:8080/article/nice"
			jsonBytes := []byte(test.reqBody)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
			req.Header.Set("Content-type", "application/json")

			res := httptest.NewRecorder()

			// ハンドラにメソッドに用意した引数を渡す
			aCon.PostArticleHandler(res, req)

			// outputと期待値との検証
			if res.Code != test.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", test.resultCode, res.Code)
			}
		})
	}
}
