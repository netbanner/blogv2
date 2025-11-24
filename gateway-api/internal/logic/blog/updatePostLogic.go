// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"blogV2/blog-rpc/blog"
	"context"
	"fmt"
	"strconv"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdatePostLogic 修改文章
func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostLogic) UpdatePost(req *types.AddPostInfoReq) error {
	// 从上下文中获取文章ID（通过路径参数传递）
	idVal := l.ctx.Value("id")
	id, ok := idVal.(string)
	if !ok {
		return fmt.Errorf("failed to get post id from context")
	}

	// 转换为int64类型
	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.BlogRPC.UpdatePost(l.ctx, &blog.UpdatePostReq{
		Id:      postId,
		Content: req.Content,
		Title:   req.Title,
	})
	if err != nil {
		return err
	}

	return nil
}
