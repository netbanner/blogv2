package svc

import (
	"blogV2/blog-rpc/internal/config"
	"blogV2/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentsModel
	PostModel    model.PostsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentsModel(dbConn, c.Cache),
		PostModel:    model.NewPostsModel(dbConn, c.Cache),
	}
}
