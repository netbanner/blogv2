// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"blogV2/user-rpc/usercenter"
	"context"
	"fmt"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {

	logx.Info("获取用户信息")
	getUserInfoResp, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("用户信息 %v\n", getUserInfoResp.Id)
	return &types.GetUserInfoResp{
		Id:   getUserInfoResp.Id,
		Name: getUserInfoResp.Name,
	}, nil
}
