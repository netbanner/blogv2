package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCurrentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentPostLogic {
	return &GetCurrentPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询当前用户的文章
func (l *GetCurrentPostLogic) GetCurrentPost(in *blog.GetCurrentPostReq) (*blog.GetCurrentPostResp, error) {
	// todo: add your logic here and delete this line

	return &blog.GetCurrentPostResp{}, nil
}
