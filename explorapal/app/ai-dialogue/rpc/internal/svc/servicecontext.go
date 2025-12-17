package svc

import (
	"explorapal/app/ai-dialogue/rpc/internal/config"
	"explorapal/third/openai"
)

type ServiceContext struct {
	Config config.Config

	// AI服务客户端
	AIClient *openai.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化AI客户端
	aiConfig := &openai.Config{
		TAL_MLOPS_APP_ID:  c.AIService.TAL_MLOPS_APP_ID,
		TAL_MLOPS_APP_KEY: c.AIService.TAL_MLOPS_APP_KEY,
		BaseURL:           c.AIService.BaseURL,
		Timeout:           c.AIService.Timeout,
		MaxTokens:         c.AIService.MaxTokens,
		Temperature:       c.AIService.Temperature,
	}
	aiClient := openai.NewClient(aiConfig)

	return &ServiceContext{
		Config:   c,
		AIClient: aiClient,
	}
}
