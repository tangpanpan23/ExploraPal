package hps

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionsModel = (*customQuestionsModel)(nil)

type (
	// QuestionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionsModel.
	QuestionsModel interface {
		questionsModel
	}

	customQuestionsModel struct {
		*defaultQuestionsModel
	}
)

// NewQuestionsModel returns a model for the database table.
func NewQuestionsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuestionsModel {
	return &customQuestionsModel{
		defaultQuestionsModel: newQuestionsModel(conn, c, opts...),
	}
}
