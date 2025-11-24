// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"blogV2/blog-rpc/blog"
	"context"
	"errors"
	"fmt"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetCurrentPostLogic 查询当前用户的文章
func NewGetCurrentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentPostLogic {
	return &GetCurrentPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentPostLogic) GetCurrentPost() (resp []types.PostInfo, err error) {
	userIdVal := l.ctx.Value("userId")
	userId, ok := userIdVal.(int64)
	if !ok {
		return nil, fmt.Errorf("failed to get userId from context")
	}
	postResp, err := l.svcCtx.BlogRPC.GetCurrentPost(l.ctx, &blog.GetCurrentPostReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.New("查询文章失败")
	}
	resp = make([]types.PostInfo, len(postResp.Posts))
	for i, postInfo := range postResp.Posts {
		resp[i] = types.PostInfo{
			ID:        postInfo.Id,
			Title:     postInfo.Title,
			Content:   postInfo.Content,
			UserID:    postInfo.UserId,
			CmtStatus: postInfo.CmtStatus,
		}
	}
	return resp, nil
}
