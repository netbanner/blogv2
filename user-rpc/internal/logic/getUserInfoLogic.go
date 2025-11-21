package logic

import (
	"context"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/user-center"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *user_center.GetUserInfoReq) (*user_center.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user_center.GetUserInfoResp{}, nil
}
