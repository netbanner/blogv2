// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"blogV2/gateway-api/internal/config"
	"blogV2/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest/token"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	CommentsModel model.CommentsModel
	UsersModel    model.UsersModel
	PostsModel    model.PostsModel
	Auth          *token.TokenParser
}

// 正确初始化 TokenParser 的方式（根据 go-zero v1.9.2 实现）
func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DB.DataSource)

	// 直接传入 secret 和 expire（如果构造函数支持）
	authParser := token.NewTokenParser()

	return &ServiceContext{
		Config:        c,
		UsersModel:    model.NewUsersModel(dbConn, c.Cache),
		CommentsModel: model.NewCommentsModel(dbConn, c.Cache),
		PostsModel:    model.NewPostsModel(dbConn, c.Cache),
		Auth:          authParser,
	}
}

// 如果你确实需要一个自定义 Auth 类型来封装 Token 生成逻辑，可以这样定义：
type Auth struct {
	AccessSecret string
	AccessExpire int64
}

func (a *Auth) GenerateToken(now time.Time, secret string, expire int64) (string, error) {
	claims := jwt.MapClaims{
		"exp": now.Add(time.Second * time.Duration(expire)).Unix(),
		"iat": now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
