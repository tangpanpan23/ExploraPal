package hps

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ObservationsModel = (*customObservationsModel)(nil)

type (
	// ObservationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customObservationsModel.
	ObservationsModel interface {
		observationsModel
	}

	customObservationsModel struct {
		*defaultObservationsModel
	}
)

// NewObservationsModel returns a model for the database table.
func NewObservationsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ObservationsModel {
	return &customObservationsModel{
		defaultObservationsModel: newObservationsModel(conn, c, opts...),
	}
}
