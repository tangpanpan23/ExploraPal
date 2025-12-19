package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 统计配置 - 慢调用阈值
	StatConf zrpc.StatConf

	// 视频处理配置
	SystemVideo struct {
		Enabled           bool     `json:",default=true"`
		MaxDuration       int      `json:",default=300"`
		SupportedFormats  []string `json:",default=["mp4","avi","mov","webm"]"`
		DefaultResolution string   `json:",default=1920x1080"`
	}

	// AI服务配置 (TAL MLOps平台)
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
