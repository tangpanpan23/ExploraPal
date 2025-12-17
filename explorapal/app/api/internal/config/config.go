package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// JWT配置
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	// 数据库配置
	DBConfig struct {
		DataSource string
	}

	// 缓存配置
	Cache cache.CacheConf

	// 内部AI服务配置 (TAL MLOps平台)
	AIService struct {
		TAL_MLOPS_APP_ID  string
		TAL_MLOPS_APP_KEY string
		BaseURL           string
		Timeout           int
		MaxTokens         int
		Temperature       float32
	}

	// CORS配置
	CORS struct {
		AllowOrigins     []string
		AllowMethods     []string
		AllowHeaders     []string
		AllowCredentials bool
	}
}
