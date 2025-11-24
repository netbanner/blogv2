package logic

import (
	"blogV2/model"
	"context"
	"database/sql"
	"errors"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostCommentLogic {
	return &AddPostCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddPostComment 添加评论
func (l *AddPostCommentLogic) AddPostComment(in *blog.AddPostCommentReq) (*blog.AddPostCommentResp, error) {

	post, err := l.svcCtx.PostModel.FindOne(l.ctx, uint64(in.PostId))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}
	if post == nil {
		return nil, errors.New("文章不存在")
	}

	comment := &model.Comments{
		Content: sql.NullString{String: in.Content, Valid: true},
		PostId:  sql.NullInt64{Int64: in.PostId, Valid: true},
		UserId:  sql.NullInt64{Int64: in.UserId, Valid: true},
	}
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		return nil, err
	}

	return &blog.AddPostCommentResp{Success: true, Message: "添加评论成功"}, nil
}
