package hps

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProjectActivitiesModel = (*customProjectActivitiesModel)(nil)

type (
	// ProjectActivitiesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectActivitiesModel.
	ProjectActivitiesModel interface {
		projectActivitiesModel
	}

	customProjectActivitiesModel struct {
		*defaultProjectActivitiesModel
	}
)

// NewProjectActivitiesModel returns a model for the database table.
func NewProjectActivitiesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProjectActivitiesModel {
	return &customProjectActivitiesModel{
		defaultProjectActivitiesModel: newProjectActivitiesModel(conn, c, opts...),
	}
}
