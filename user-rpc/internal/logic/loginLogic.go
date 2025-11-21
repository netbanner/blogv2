package logic

import (
	"context"
	"fmt"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/user-center"

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

// 用户登录
func (l *LoginLogic) Login(in *user_center.LoginReq) (*user_center.LoginResp, error) {
	// todo: add your logic here and delete this line
	fmt.Println("登录测试")
	return &user_center.LoginResp{}, nil
}
