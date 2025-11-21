// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewLoginLogic 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}
	user, err := l.svcCtx.UsersModel.FindByUserName(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if !user.Password.Valid || user.Password.String != req.Password {
		return nil, errors.New("密码错误")
	}
	now := time.Now()
	expireDuration := time.Duration(l.svcCtx.Config.Auth.AccessExpire) * time.Second
	expireAt := now.Add(expireDuration).Unix()

	// 使用 golang-jwt/jwt/v4 生成 token
	claims := jwt.MapClaims{
		"userId":   user.Id,
		"username": user.UserName,
		"exp":      expireAt,
		"iat":      now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{Token: signedToken, ExpireAt: expireAt}, nil
}
