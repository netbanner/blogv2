package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByUserIdLogic {
	return &GetPostByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPostByUserId 根据用户id查询所有文章 查询文章，TODO 也需把文章对应的评论也查询处理
func (l *GetPostByUserIdLogic) GetPostByUserId(in *blog.GetPostByUserIdReq) (*blog.GetPostByUserIdResp, error) {
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

	return &blog.GetPostByUserIdResp{
		Posts: postsInfo,
	}, nil

}
