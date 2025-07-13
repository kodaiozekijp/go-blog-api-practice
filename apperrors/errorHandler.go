package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
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
