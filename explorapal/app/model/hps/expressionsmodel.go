package hps

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExpressionsModel = (*customExpressionsModel)(nil)

type (
	// ExpressionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customExpressionsModel.
	ExpressionsModel interface {
		expressionsModel
	}

	customExpressionsModel struct {
		*defaultExpressionsModel
	}
)

// NewExpressionsModel returns a model for the database table.
func NewExpressionsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ExpressionsModel {
	return &customExpressionsModel{
		defaultExpressionsModel: newExpressionsModel(conn, c, opts...),
	}
}
