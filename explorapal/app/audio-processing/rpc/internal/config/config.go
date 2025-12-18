package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 统计配置 - 慢调用阈值
	StatConf zrpc.StatConf

	// 语音处理配置
	SystemTTS struct {
		Enabled bool    `json:",default=true"`
		Voice   string  `json:",default=zh-CN-XiaoxiaoNeural"`
		Speed   float64 `json:",default=1.0"`
	}
	SystemSTT struct {
		Enabled  bool   `json:",default=true"`
		Language string `json:",default=zh-CN"`
	}

	// AI服务配置 (TAL MLOps平台)
	AIService struct {
		TAL_MLOPS_APP_ID  string
		TAL_MLOPS_APP_KEY string
		BaseURL           string
		Timeout           int
	}

	// 缓存配置
	Cache cache.CacheConf
}
