package model

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	// BaseModel 基础模型接口
	BaseModel interface {
		Insert(ctx context.Context, data interface{}) (sql.Result, error)
		FindOne(ctx context.Context, id interface{}) (interface{}, error)
		Update(ctx context.Context, data interface{}) error
		Delete(ctx context.Context, id interface{}) error
	}

	// CacheConfig 缓存配置
	CacheConfig struct {
		Cache cache.CacheConf
	}

	// DBConfig 数据库配置
	DBConfig struct {
		DSN string
	}
)