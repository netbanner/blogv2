package logic

import (
	"blogV2/model"
	"context"
	"errors"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/usercenter"

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

// GetUserInfo 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *usercenter.GetUserInfoReq) (*usercenter.GetUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("用户不存在")
		} else {
			// 其他数据库错误
			return nil, err
		}
		return nil, err
	}

	return &usercenter.GetUserInfoResp{
		Id:   int64(user.Id),
		Name: user.UserName,
	}, nil
}
