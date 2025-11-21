package logic

import (
	"context"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/user-center"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户信息
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user_center.UpdateUserInfoReq) (*user_center.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user_center.UpdateUserInfoResp{}, nil
}
