package hps

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AchievementsModel = (*customAchievementsModel)(nil)

type (
	// AchievementsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAchievementsModel.
	AchievementsModel interface {
		achievementsModel
	}

	customAchievementsModel struct {
		*defaultAchievementsModel
	}
)

// NewAchievementsModel returns a model for the database table.
func NewAchievementsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AchievementsModel {
	return &customAchievementsModel{
		defaultAchievementsModel: newAchievementsModel(conn, c, opts...),
	}
}
