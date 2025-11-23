// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"blogV2/user-rpc/usercenter"
	"context"
	"errors"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	logx.Info("用户登录")
	loginResp, err := l.svcCtx.UserRPC.Login(l.ctx, &usercenter.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.New("登录错误")
	}
	return &types.LoginResp{
		Id:       loginResp.Id,
		Token:    loginResp.Token,
		ExpireAt: loginResp.ExpireAt,
	}, nil

}
