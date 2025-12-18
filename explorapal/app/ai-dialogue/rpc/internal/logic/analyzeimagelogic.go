package logic

import (
	"context"
	"fmt"
	"strings"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/ai-dialogue/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalyzeImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalyzeImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzeImageLogic {
	return &AnalyzeImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnalyzeImageLogic) AnalyzeImage(in *aidialogue.AnalyzeImageReq) (*aidialogue.AnalyzeImageResp, error) {
	// TODO: 实现图片分析逻辑
	// 注意：需要先运行 protoc 生成 aidialogue 包
	// 命令: protoc --go_out=. --go-grpc_out=. ai-dialogue.proto
	
	result, err := l.svcCtx.AIClient.AnalyzeImage(l.ctx, in.ImageUrl, in.Prompt)
	if err != nil {
		l.Logger.Errorf("图片分析失败，使用模拟响应: %v", err)
		// 返回默认的分析结果
		return l.getDefaultImageAnalysisResult(in.ImageUrl, in.Prompt), nil
	}

	// 检查结果是否有效
	if result == nil {
		return &aidialogue.AnalyzeImageResp{
			Status: 500,
			Msg:    "图片分析失败，无法获取结果",
		}, fmt.Errorf("无法获取图片分析结果")
	}

	return &aidialogue.AnalyzeImageResp{
		Status:        200,
		Msg:           "图片分析成功",
		ObjectName:    sanitizeUTF8(result.ObjectName),
		Category:      sanitizeUTF8(result.Category),
		Confidence:    float32(result.Confidence),
		Description:   sanitizeUTF8(result.Description),
		KeyFeatures:   sanitizeUTF8Slice(result.KeyFeatures),
		ScientificName: sanitizeUTF8(result.ScientificName),
	}, nil
}

// getDefaultImageAnalysisResult 当AI服务完全不可用时，返回默认的分析结果
func (l *AnalyzeImageLogic) getDefaultImageAnalysisResult(imageURL, prompt string) *aidialogue.AnalyzeImageResp {
	// 基于URL关键词判断内容类型
	objectName := "未知物体"
	category := "general"
	description := "由于AI服务暂时不可用，这里提供一个模拟的分析结果。在实际环境中，这个结果将由AI模型生成。"
	scientificName := "未知"

	// 简单的关键词匹配
	if containsKeyword(imageURL, "dinosaur", "恐龙", "化石") {
		objectName = "恐龙化石"
		category = "dinosaur"
		description = "这看起来像是一块恐龙化石，包含了古代生物的遗骸。"
		scientificName = "恐龙类"
	} else if containsKeyword(imageURL, "rocket", "火箭", "太空") {
		objectName = "火箭模型"
		category = "rocket"
		description = "这是一枚火箭模型，用于太空探索。"
		scientificName = "火箭"
	} else if containsKeyword(imageURL, "minecraft", "方块", "游戏") {
		objectName = "Minecraft建筑"
		category = "minecraft"
		description = "这是Minecraft游戏中的建筑作品。"
		scientificName = "虚拟建筑"
	}

	return &aidialogue.AnalyzeImageResp{
		Status:        200,
		Msg:           "图片分析成功（使用模拟响应）",
		ObjectName:    sanitizeUTF8(objectName),
		Category:      sanitizeUTF8(category),
		Confidence:    0.85,
		Description:   sanitizeUTF8(description),
		KeyFeatures:   sanitizeUTF8Slice([]string{"特征分析", "形态识别", "内容描述"}),
		ScientificName: sanitizeUTF8(scientificName),
	}
}

// containsKeyword 检查字符串是否包含任一关键词
func containsKeyword(text string, keywords ...string) bool {
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(text), strings.ToLower(keyword)) {
			return true
		}
	}
	return false
}

