package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostCommentHandler(t *testing.T) {
	// テストケースを用意
	var tests = []struct {
		title      string
		reqBody    string
		resultCode int
	}{
		{title: "decoded test", reqBody: `{"article_id": 1,"comment_id": 1,"message": "ccc","created_at": "2024-01-01T00:00:00Z"}`, resultCode: http.StatusOK},
		{title: "can't be decoded test", reqBody: `{"article_id":,"message":"aaa"}`, resultCode: http.StatusBadRequest},
	}

	// テストの実行
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// ハンドラに渡す引数の用意
			url := "http://localhost:8080/comment"
			jsonBytes := []byte(test.reqBody)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
			req.Header.Set("Content-type", "application/json")

			res := httptest.NewRecorder()

			// ハンドラにメソッドに用意した引数を渡す
			cCon.PostCommentHandler(res, req)

			// outputと期待値との検証
			if res.Code != test.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", test.resultCode, res.Code)
			}
		})
	}
}
