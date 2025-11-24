package logic

import (
	"blogV2/common/util"
	"context"
	"errors"
	"strconv"
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
	expireDuration := time.Duration(l.svcCtx.Config.Jwt.AccessExpire) * time.Second
	expireAt := now.Add(expireDuration).Unix()

	// 使用标准化方法生成 token
	signedToken, err := util.GenerateToken(
		int64(user.Id),
		user.UserName,
		l.svcCtx.Config.Jwt.AccessSecret,
		expireDuration,
	)
	if err != nil {
		return nil, err
	}

	return &usercenter.LoginResp{
		Token:    signedToken,
		ExpireAt: strconv.FormatInt(expireAt, 10),
	}, nil
}
