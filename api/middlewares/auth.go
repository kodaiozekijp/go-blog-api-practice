package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/kodaiozekijp/go-blog-api-practice/apperrors"
	"google.golang.org/api/idtoken"
)

const (
	googleClientID = "629786876842-u95882tfn5c63p4ng5dtikh7cv677am1.apps.googleusercontent.com"
)

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

// トークン検証を行うミドルウェア
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ヘッダからuserizationフィールドを抜き出す
		authorization := req.Header.Get("authorization")

		// authorrizationフィールドが"Bearer [IDトークン]"の形になっているか検証
		authorizationHeaders := strings.Split(authorization, " ") //空白区切りで2つに分かれるか
		if len(authorizationHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authorizationHeaders[0], authorizationHeaders[1] // 空白区切りで分けた1つ目がBearerで、2つ目が空でないか
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// IDトークン検証
		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// nameフィールドをpayloadから抜き出す
		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// コンテキストのnameフィールドを設定
		req = SetUserName(req, name.(string))

		// ハンドラへ
		next.ServeHTTP(w, req)
	})
}
