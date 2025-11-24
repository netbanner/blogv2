package logic

import (
	"blogV2/model"
	"context"
	"database/sql"
	"fmt"
	"log"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePost 修改文章
func (l *UpdatePostLogic) UpdatePost(in *blog.UpdatePostReq) (*blog.UpdatePostResp, error) {
	// 构造更新数据
	post := &model.Posts{
		Id: uint64(in.Id),
	}

	// 只有非空字段才参与更新
	if in.Title != "" {
		post.Title = sql.NullString{String: in.Title, Valid: true}
	}
	if in.Content != "" {
		post.Content = sql.NullString{String: in.Content, Valid: true}
	}
	if in.UserId > 0 {
		post.UserId = sql.NullInt64{Int64: in.UserId, Valid: true}
	}
	err := l.svcCtx.PostModel.Update(l.ctx, post)
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to update post %d: %v", in.Id, err)
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return &blog.UpdatePostResp{Success: true, Message: "修改文章成功"}, nil
}
