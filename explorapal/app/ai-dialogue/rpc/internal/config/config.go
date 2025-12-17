package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 内部AI服务配置 (TAL MLOps平台)
	AIService struct {
		TAL_MLOPS_APP_ID  string
		TAL_MLOPS_APP_KEY string
		BaseURL           string
		Timeout           int
		MaxTokens         int
		Temperature       float32
	}

	// 缓存配置
	Cache cache.CacheConf
}
