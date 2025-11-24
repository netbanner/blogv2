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

type AddPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddPostLogic 添加文章
func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostLogic) AddPost(req *types.AddPostInfoReq) error {
	userIdVal := l.ctx.Value("userId")
	userId, ok := userIdVal.(int64)
	if !ok {
		return errors.New("用户id不存在")
	}
	_, err := l.svcCtx.BlogRPC.AddPost(l.ctx, &blog.AddPostReq{
		Content: req.Content,
		Title:   req.Title,
		UserId:  userId,
	})
	if err != nil {
		return errors.New("添加失败")
	}

	return nil
}
