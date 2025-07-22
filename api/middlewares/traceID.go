package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1 // トレースIDを1からの連番で保存
	mu    sync.Mutex
)

// トレースIDを設定する際に使用する構造体
type traceIDKey struct{}

// トレースIDを返却する
func newTraceID() int {
	var no int // トレースIDの一時保存用

	mu.Lock()   // ロックを取得する
	no = logNo  // トレースIDをno変数にコピー
	logNo += 1  // トレースIDを+1する
	mu.Unlock() // ロックを解除する

	return no
}

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
