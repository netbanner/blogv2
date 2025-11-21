package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentsLogic {
	return &GetPostCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文章的所有评论列表
func (l *GetPostCommentsLogic) GetPostComments(in *blog.GetPostCommentsReq) (*blog.GetPostCommentsResp, error) {
	// todo: add your logic here and delete this line

	return &blog.GetPostCommentsResp{}, nil
}
