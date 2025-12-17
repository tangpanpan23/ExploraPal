package logic

import (
	"context"

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
		l.Logger.Errorf("图片分析失败: %v", err)
		return &aidialogue.AnalyzeImageResp{
			Status: 500,
			Msg:    "图片分析失败",
		}, err
	}

	return &aidialogue.AnalyzeImageResp{
		Status:        200,
		Msg:           "图片分析成功",
		ObjectName:    result.ObjectName,
		Category:      result.Category,
		Confidence:    float32(result.Confidence),
		Description:   result.Description,
		KeyFeatures:   result.KeyFeatures,
		ScientificName: result.ScientificName,
	}, nil
}
