package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostCommentLogic {
	return &AddPostCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加评论
func (l *AddPostCommentLogic) AddPostComment(in *blog.AddPostCommentReq) (*blog.AddPostCommentResp, error) {
	// todo: add your logic here and delete this line

	return &blog.AddPostCommentResp{}, nil
}
