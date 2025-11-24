package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPostById 根据文章id查询文章
func (l *GetPostByIdLogic) GetPostById(in *blog.GetPostByIdReq) (*blog.PostInfo, error) {
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}

	return &blog.PostInfo{
		Id:        int64(post.Id),
		Title:     post.Title.String,
		Content:   post.Content.String,
		UserId:    post.UserId.Int64,
		CmtStatus: post.CmtStatus,
	}, nil
}
