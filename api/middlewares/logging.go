package middlewares

import (
	"log"
	"net/http"
)

func LoggingMiddleWare(next http.Handler) http.Handler {
	// http.Handlerを返却する
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println(req.RequestURI, req.Method)
		// リクエストを処理
		next.ServeHTTP(w, req)
	})
}
