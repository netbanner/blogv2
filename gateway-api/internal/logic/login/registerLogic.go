// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
