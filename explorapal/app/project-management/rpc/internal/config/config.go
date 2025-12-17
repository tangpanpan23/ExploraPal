package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 数据库配置
	DBConfig struct {
		DataSource string
	}

	// 缓存配置
	Cache cache.CacheConf
}
