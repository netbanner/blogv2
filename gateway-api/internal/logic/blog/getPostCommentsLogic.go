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

type GetPostCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetPostCommentsLogic 获取文章的所有评论列表
func NewGetPostCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentsLogic {
	return &GetPostCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostCommentsLogic) GetPostComments(req *types.GetPostById) (resp []types.CommentInfo, err error) {
	commentsResp, err := l.svcCtx.BlogRPC.GetPostComments(l.ctx, &blog.GetPostCommentsReq{
		PostId: req.ID,
	})
	if err != nil {
		return nil, errors.New("根据文章id查询评论失败")
	}
	resp = make([]types.CommentInfo, len(commentsResp.Comments))
	for i, comment := range commentsResp.Comments {
		resp[i] = types.CommentInfo{
			ID:      comment.Id,
			Content: comment.Content,
			UserID:  uint(comment.UserId),
		}
	}

	return resp, nil
}
