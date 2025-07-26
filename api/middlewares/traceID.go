package middlewares

import (
	"sync"
)

var (
	logNo int = 1 // トレースIDを1からの連番で保存
	mu    sync.Mutex
)

// トレースIDを返却する
func newTraceID() int {
	var no int // トレースIDの一時保存用

	mu.Lock()   // ロックを取得する
	no = logNo  // トレースIDをno変数にコピー
	logNo += 1  // トレースIDを+1する
	mu.Unlock() // ロックを解除する

	return no
}
