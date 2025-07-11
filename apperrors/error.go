package apperrors

type MyAppError struct {
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  // エラーチェーンの為の内部エラー
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) UnWrap() error {
	return myErr.Err
}
