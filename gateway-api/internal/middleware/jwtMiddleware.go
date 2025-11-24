package middleware

import (
	"blogV2/common/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type JwtMiddleware struct {
	secret string
}

func NewJwtMiddleware(secret string) *JwtMiddleware {
	return &JwtMiddleware{
		secret: secret,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 从 Authorization header 中提取 token
		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			httpx.Error(writer, fmt.Errorf("missing authorization header"))
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			httpx.Error(writer, fmt.Errorf("invalid authorization header format"))
			return
		}
		chaims, err := util.ParseToken(tokenString, m.secret)
		if err != nil {
			httpx.Error(writer, fmt.Errorf("invalid token: %v", err))
			return
		}

		//将用户信息添加到context 中
		ctx := context.WithValue(request.Context(), "userId", chaims.UserId)
		ctx = context.WithValue(ctx, "userName", chaims.UserName)
		//继续处理请求
		next(writer, request.WithContext(ctx))
	}

}
