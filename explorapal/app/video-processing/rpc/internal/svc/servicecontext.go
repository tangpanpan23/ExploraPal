package svc

import (
	"explorapal/app/video-processing/rpc/internal/config"
	"explorapal/third/openai"

	"github.com/zeromicro/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config   config.Config
	Cache    cache.CacheConf
	AIClient *openai.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cache:  c.Cache,
		AIClient: openai.NewClient(&openai.Config{
			TAL_MLOPS_APP_ID:  c.AIService.TAL_MLOPS_APP_ID,
			TAL_MLOPS_APP_KEY: c.AIService.TAL_MLOPS_APP_KEY,
			BaseURL:           c.AIService.BaseURL,
			Timeout:           c.AIService.Timeout,
			MaxTokens:         c.AIService.MaxTokens,
			Temperature:       c.AIService.Temperature,
		}),
	}
}
