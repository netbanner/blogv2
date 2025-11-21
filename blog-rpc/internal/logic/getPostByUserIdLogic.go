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

// 根据用户id查询所有文章
func (l *GetPostByUserIdLogic) GetPostByUserId(in *blog.GetPostByUserIdReq) (*blog.GetPostByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &blog.GetPostByUserIdResp{}, nil
}
