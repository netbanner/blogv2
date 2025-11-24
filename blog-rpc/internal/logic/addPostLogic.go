package logic

import (
	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"
	"blogV2/model"
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddPost 添加文章
func (l *AddPostLogic) AddPost(in *blog.AddPostReq) (*blog.AddPostResp, error) {
	logx.Info("添加文章")
	_, err := l.svcCtx.PostModel.Insert(l.ctx, &model.Posts{
		UserId:  sql.NullInt64{Int64: in.UserId, Valid: true},
		Title:   sql.NullString{String: in.Title, Valid: true},
		Content: sql.NullString{String: in.Content, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &blog.AddPostResp{Success: true, Message: "添加文章成功"}, nil
}
