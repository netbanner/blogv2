package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加文章
func (l *AddPostLogic) AddPost(in *blog.AddPostReq) (*blog.AddPostResp, error) {
	// todo: add your logic here and delete this line

	return &blog.AddPostResp{}, nil
}
