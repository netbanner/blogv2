package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentsModel = (*customCommentsModel)(nil)

type (
	// CommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentsModel.
	CommentsModel interface {
		commentsModel
		FindAll(ctx context.Context, id int64) ([]Comments, error)
	}

	customCommentsModel struct {
		*defaultCommentsModel
	}
)

// NewCommentsModel returns a model for the database table.
func NewCommentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentsModel {
	return &customCommentsModel{
		defaultCommentsModel: newCommentsModel(conn, c, opts...),
	}
}

func (m *customCommentsModel) FindAll(ctx context.Context, id int64) ([]Comments, error) {
	var resp []Comments
	err := m.QueryRowCtx(ctx, &resp, "", func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `post_id` = ? order by id desc ", commentsRows, m.table)
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
