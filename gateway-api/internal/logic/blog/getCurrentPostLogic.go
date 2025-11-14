// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"context"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询当前用户的文章
func NewGetCurrentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentPostLogic {
	return &GetCurrentPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentPostLogic) GetCurrentPost() (resp []types.PostInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
