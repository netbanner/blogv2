package logic

import (
	"context"
	"errors"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentsLogic {
	return &GetPostCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPostComments 获取文章的所有评论列表
func (l *GetPostCommentsLogic) GetPostComments(in *blog.GetPostCommentsReq) (*blog.GetPostCommentsResp, error) {

	comments, err := l.svcCtx.CommentModel.FindAll(l.ctx, in.PostId)
	if err != nil {
		return nil, errors.New("查询评论错误")
	}
	var commentList []*blog.CommentInfo
	for _, comment := range comments {
		commentList = append(commentList, &blog.CommentInfo{
			Id:     int64(comment.Id),
			PostId: comment.PostId.Int64,
			UserId: uint32(comment.UserId.Int64),
			Content: func() string {
				if comment.Content.Valid {
					return comment.Content.String
				}
				return ""
			}(),
		})
	}

	return &blog.GetPostCommentsResp{
		Comments: commentList,
	}, nil
}
