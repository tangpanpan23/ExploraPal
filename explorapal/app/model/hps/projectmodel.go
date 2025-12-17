package hps

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProjectModel = (*customProjectModel)(nil)

type (
	// ProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectModel.
	ProjectModel interface {
		projectModel
		FindByUserID(ctx context.Context, userID int64, page, pageSize int64) ([]*Project, error)
		FindByCategory(ctx context.Context, userID int64, category string, page, pageSize int64) ([]*Project, error)
		UpdateProgress(ctx context.Context, projectID int64, progress int32) error
		UpdateLastActivity(ctx context.Context, projectID int64) error
	}

	customProjectModel struct {
		*defaultProjectModel
	}
)

// NewProjectModel returns a model for the database table.
func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectModel {
	return &customProjectModel{
		defaultProjectModel: newProjectModel(conn, c),
	}
}

func (m *customProjectModel) FindByUserID(ctx context.Context, userID int64, page, pageSize int64) ([]*Project, error) {
	var resp []*Project
	offset := (page - 1) * pageSize

	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `delete_time` IS NULL order by `create_time` desc limit ?,?",
		projectRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userID, offset, pageSize)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customProjectModel) FindByCategory(ctx context.Context, userID int64, category string, page, pageSize int64) ([]*Project, error) {
	var resp []*Project
	offset := (page - 1) * pageSize

	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `category` = ? and `delete_time` IS NULL order by `create_time` desc limit ?,?",
		projectRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userID, category, offset, pageSize)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customProjectModel) UpdateProgress(ctx context.Context, projectID int64, progress int32) error {
	projectProjectIdKey := fmt.Sprintf("%s%v", cacheProjectProjectIdPrefix, projectID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `progress` = ? where `project_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, progress, projectID)
	}, projectProjectIdKey)
	return err
}

func (m *customProjectModel) UpdateLastActivity(ctx context.Context, projectID int64) error {
	projectProjectIdKey := fmt.Sprintf("%s%v", cacheProjectProjectIdPrefix, projectID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `last_activity_at` = now() where `project_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, projectID)
	}, projectProjectIdKey)
	return err
}
