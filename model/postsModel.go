package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostsModel = (*customPostsModel)(nil)

type (
	// PostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsModel.
	PostsModel interface {
		postsModel
		FindPostByUserId(ctx context.Context, id uint64) ([]Posts, error)
	}

	customPostsModel struct {
		*defaultPostsModel
	}
)

// NewPostsModel returns a model for the database table.
func NewPostsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostsModel {
	return &customPostsModel{
		defaultPostsModel: newPostsModel(conn, c, opts...),
	}
}

func (m *defaultPostsModel) FindPostByUserId(ctx context.Context, id uint64) ([]Posts, error) {

	var resp []Posts
	err := m.QueryRowCtx(ctx, &resp, "", func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? order by id desc ", postsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
