package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据文章id查询文章
func (l *GetPostByIdLogic) GetPostById(in *blog.GetPostByIdReq) (*blog.PostInfo, error) {
	// todo: add your logic here and delete this line

	return &blog.PostInfo{}, nil
}
