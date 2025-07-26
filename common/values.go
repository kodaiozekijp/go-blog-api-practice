package common

import (
	"context"
	"net/http"
)

// トレースIDを設定する際に使用する構造体
type traceIDKey struct{}

// コンテキストにトレースIDを付加して返却する
func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// コンテキストからトレースIDを取得して返却する
func GetTraceID(ctx context.Context) int {
	traceIDAny := ctx.Value(traceIDKey{})
	// int型にキャストする
	if traceID, ok := traceIDAny.(int); ok {
		return traceID
	}
	return 0
}

// コンテキストの中のnameフィールドに対応させるキー構造体
type userNameKey struct{}

// コンテキストからnameフィールドの値を取得する
func GetUserName(ctx context.Context) string {
	id := ctx.Value(userNameKey{})

	// コンテキストから取得したnameフィールドの値をstringにして返却する
	if userNameStr, ok := id.(string); ok {
		return userNameStr
	}
	return ""
}

// コンテキストにnameフィールドの値をセットする
func SetUserName(req *http.Request, name string) *http.Request {
	ctx := req.Context() // 引数のリクエストからコンテキストを取得

	ctx = context.WithValue(ctx, userNameKey{}, name) // コンテキストのnameフィールドの値を設定
	req = req.WithContext(ctx)                        // リクエストにコンテキストを設定

	return req
}
