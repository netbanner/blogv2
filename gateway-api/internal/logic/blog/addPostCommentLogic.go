// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"context"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加评论
func NewAddPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostCommentLogic {
	return &AddPostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostCommentLogic) AddPostComment(req *types.AddPostComment) error {
	// todo: add your logic here and delete this line

	return nil
}
