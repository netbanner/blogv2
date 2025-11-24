// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"blogV2/blog-rpc/blog"
	"context"
	"errors"
	"fmt"
	"strconv"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetPostByIdLogic 根据文章id查询文章
func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByIdLogic) GetPostById() (resp *types.PostInfo, err error) {
	// 从上下文中获取文章ID（通过路径参数传递）
	idVal := l.ctx.Value("id")
	id, ok := idVal.(string)
	if !ok {
		return nil, fmt.Errorf("failed to get post id from context")
	}

	// 转换为int64类型
	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid post id")
	}
	postResp, err := l.svcCtx.BlogRPC.GetPostById(l.ctx, &blog.GetPostByIdReq{
		Id: postId,
	})
	if err != nil {
		return nil, errors.New("查询文章失败")
	}
	resp = &types.PostInfo{
		ID:        postResp.Id,
		Title:     postResp.Title,
		Content:   postResp.Content,
		UserID:    postResp.UserId,
		CmtStatus: postResp.CmtStatus,
	}

	return resp, nil
}
