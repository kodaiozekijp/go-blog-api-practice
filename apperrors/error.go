package apperrors

type MyAppError struct {
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  // エラーチェーンの為の内部エラー
}

// エラー情報出力時に表示する内容を返却する
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// Errフィールドに格納された内部エラーを返却する
func (myErr *MyAppError) UnWrap() error {
	return myErr.Err
}
