package logic

import (
	"context"

	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeletePost 删除文章
func (l *DeletePostLogic) DeletePost(in *blog.DeletePostReq) (*blog.DeletePostResp, error) {
	logx.Info("删除文章")
	err := l.svcCtx.PostModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}

	return &blog.DeletePostResp{Success: true, Message: "删除文章成功"}, nil
}
