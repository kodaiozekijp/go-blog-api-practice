package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/kodaiozekijp/go-blog-api-practice/api/middlewares"
)

// エラーの内容に応じた
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// 受け渡されたerrがMyAppError型か確認する
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	// トレースIDを使用しロギングする
	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	// 元となったエラーに応じたステータスコードを設定する
	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParameter:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	// jsonのエンコーダーを使用し、エラーをエンコードする
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
