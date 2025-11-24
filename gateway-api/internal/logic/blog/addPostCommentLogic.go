// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"blogV2/blog-rpc/blog"
	"context"
	"errors"

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
	userIdVal := l.ctx.Value("userId")
	userId, ok := userIdVal.(int64)
	if !ok {
		return errors.New("用户id不存在")
	}
	_, err := l.svcCtx.BlogRPC.AddPostComment(l.ctx, &blog.AddPostCommentReq{
		Content: req.Content,
		PostId:  req.PostID,
		UserId:  userId,
	})

	if err != nil {
		return errors.New("添加失败")
	}

	return nil
}
