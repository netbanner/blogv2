package logic

import (
	"blogV2/model"
	"context"
	"errors"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/usercenter"

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

// UpdateUserInfo 更新用户信息
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *usercenter.UpdateUserInfoReq) (*usercenter.UpdateUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("用户已存在")
	}
	l.svcCtx.UserModel.Update(l.ctx, &model.Users{
		UserName: in.Name,
	})

	return &usercenter.UpdateUserInfoResp{Success: true, Message: "修改成功"}, nil
}
