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

type GetPostByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetPostByUserIdLogic 根据用户id查询所以文章
func NewGetPostByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByUserIdLogic {
	return &GetPostByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByUserIdLogic) GetPostByUserId(req *types.GetPostInfoReq) (resp []types.PostInfo, err error) {
	postResp, err := l.svcCtx.BlogRPC.GetPostByUserId(l.ctx, &blog.GetPostByUserIdReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errors.New("根据用户ID查询文章列表失败")
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
