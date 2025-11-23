package logic

import (
	"blogV2/model"
	"context"
	"errors"

	"blogV2/user-rpc/internal/svc"
	"blogV2/user-rpc/usercenter"

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

// Register 用户注册
func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.RegisterResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("用户已存在")
	}
	l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		UserName: in.Username,
		Password: in.Password,
		Email:    in.Email,
	})

	return &usercenter.RegisterResp{Code: 200, Msg: "注册成功"}, nil
}
