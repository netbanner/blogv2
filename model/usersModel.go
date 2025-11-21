package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		FindByUserName(ctx context.Context, userName string) (*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}
func (m *customUsersModel) FindByUserName(ctx context.Context, userName string) (*Users, error) {
	var resp Users
	query := "SELECT * FROM " + m.table + " WHERE user_name = ? LIMIT 1"
	//cacheKey := "user:name:" + userName // 构造缓存键
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &resp, query, userName)
	if err != nil {
		return nil, err
	}

	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
