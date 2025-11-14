// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"context"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据用户id查询所以文章
func NewGetPostByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByUserIdLogic {
	return &GetPostByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByUserIdLogic) GetPostByUserId(req *types.GetPostInfoReq) (resp []types.PostInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
