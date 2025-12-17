package hps

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProjectsModel = (*customProjectsModel)(nil)

type (
	// ProjectsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectsModel.
	ProjectsModel interface {
		projectsModel
		FindByUserID(ctx context.Context, userID int64, page, pageSize int64) ([]*Projects, error)
		FindByCategory(ctx context.Context, userID int64, category string, page, pageSize int64) ([]*Projects, error)
		UpdateProgress(ctx context.Context, projectID int64, progress int32) error
		UpdateLastActivity(ctx context.Context, projectID int64) error
	}

	customProjectsModel struct {
		*defaultProjectsModel
	}
)

// NewProjectsModel returns a model for the database table.
func NewProjectsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProjectsModel {
	return &customProjectsModel{
		defaultProjectsModel: newProjectsModel(conn, c, opts...),
	}
}

func (m *customProjectsModel) FindByUserID(ctx context.Context, userID int64, page, pageSize int64) ([]*Projects, error) {
	var resp []*Projects
	offset := (page - 1) * pageSize

	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `delete_time` IS NULL order by `create_time` desc limit ?,?",
		projectsRows, m.table)
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

func (m *customProjectsModel) FindByCategory(ctx context.Context, userID int64, category string, page, pageSize int64) ([]*Projects, error) {
	var resp []*Projects
	offset := (page - 1) * pageSize

	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `category` = ? and `delete_time` IS NULL order by `create_time` desc limit ?,?",
		projectsRows, m.table)
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

func (m *customProjectsModel) UpdateProgress(ctx context.Context, projectID int64, progress int32) error {
	projectsProjectIdKey := fmt.Sprintf("%s%v", cacheProjectsProjectIdPrefix, projectID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `progress` = ? where `project_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, progress, projectID)
	}, projectsProjectIdKey)
	return err
}

func (m *customProjectsModel) UpdateLastActivity(ctx context.Context, projectID int64) error {
	projectsProjectIdKey := fmt.Sprintf("%s%v", cacheProjectsProjectIdPrefix, projectID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `last_activity_at` = now() where `project_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, projectID)
	}, projectsProjectIdKey)
	return err
}
