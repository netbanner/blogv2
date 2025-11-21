package logic

import (
	"context"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/user-center"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *user_center.RegisterReq) (*user_center.RegisterResp, error) {
	// todo: add your logic here and delete this line

	return &user_center.RegisterResp{}, nil
}
