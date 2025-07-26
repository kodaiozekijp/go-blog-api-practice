package middlewares

import (
	"log"
	"net/http"

	"github.com/kodaiozekijp/go-blog-api-practice/common"
)

// 自作のResponseWriter
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

// コンストラクタを作る
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderメソッドのオーバーライド
func (rlw *resLoggingWriter) WriteHeader(code int) {
	rlw.code = code
	rlw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleWare(next http.Handler) http.Handler {
	// http.Handlerを返却する
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// トレースIDを取得する
		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		// リクエストのコンテキストをトレースID入りのコンテキストにする
		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)

		// 自作のResponseWriterを生成
		rlw := NewResLoggingWriter(w)

		// リクエストを処理
		next.ServeHTTP(rlw, req)

		// 自作ResponseWriterからロギングしたいデータを取得する
		log.Printf("[%d]response-code: %d", traceID, rlw.code)

	})
}
