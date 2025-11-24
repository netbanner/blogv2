package logic

import (
	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCurrentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentPostLogic {
	return &GetCurrentPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询当前用户的文章
func (l *GetCurrentPostLogic) GetCurrentPost(in *blog.GetCurrentPostReq) (*blog.GetCurrentPostResp, error) {

	posts, err := l.svcCtx.PostModel.FindPostByUserId(l.ctx, uint64(in.UserId))
	if err != nil {
		return nil, err
	}
	postsInfo := make([]*blog.PostInfo, len(posts))
	for i, post := range posts {
		postsInfo[i] = &blog.PostInfo{
			Id: int64(post.Id),
			//处理空值
			Title: func() string {
				if post.Title.Valid {
					return post.Title.String
				}
				return ""
			}(),
			Content: func() string {
				if post.Content.Valid {
					return post.Content.String
				}
				return ""
			}(),
		}
	}

	return &blog.GetCurrentPostResp{
		Posts: postsInfo,
	}, nil
}
