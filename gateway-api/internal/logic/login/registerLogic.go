// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"blogV2/user-rpc/userCenterClient"
	"context"
	"errors"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.ResisterResp, err error) {
	logx.Info("用户注册")
	registerResp, err := l.svcCtx.UserRPC.Register(l.ctx, &userCenterClient.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		return nil, errors.New("注册错误")
	}
	return &types.ResisterResp{
		Code: registerResp.Code,
		Msg:  registerResp.Msg,
	}, nil
}
