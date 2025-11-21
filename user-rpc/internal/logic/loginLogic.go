package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 用户登录
func (l *LoginLogic) Login(in *usercenter.LoginReq) (*usercenter.LoginResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != in.Password {
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

	return &usercenter.LoginResp{
		Token:    signedToken,
		ExpireAt: string(expireAt),
	}, nil
}
