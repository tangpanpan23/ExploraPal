package logic

import (
	"context"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type RecognizeImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecognizeImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecognizeImageLogic {
	return &RecognizeImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecognizeImageLogic) RecognizeImage(req *types.RecognizeImageReq) (resp *types.RecognizeImageResp, err error) {
	// 调用AI对话RPC服务进行图像分析
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9002"}, // AI对话RPC服务地址
		Timeout:   75000,                      // 75秒超时，与AI分析超时保持一致
	})
	if err != nil {
		l.Logger.Errorf("创建AI RPC客户端失败: %v", err)
		return nil, err
	}

	aiClient := aidialogue.NewAIDialogueServiceClient(client.Conn())

	prompt := req.Prompt
	if prompt == "" {
		prompt = "请分析这张图片，识别其中的主要物体，描述其特征和科学信息。适合儿童学习的友好方式。"
	}

	category := req.Category
	if category == "" {
		category = "general"
	}

	rpcResp, err := aiClient.AnalyzeImage(l.ctx, &aidialogue.AnalyzeImageReq{
		ImageUrl: req.ImageUrl,
		Prompt:   prompt,
		Category: category,
	})
	if err != nil {
		l.Logger.Errorf("调用AI图像分析服务失败: %v", err)
		return nil, err
	}

	resp = &types.RecognizeImageResp{
		ObjectName:    rpcResp.ObjectName,
		Category:      rpcResp.Category,
		Confidence:    rpcResp.Confidence,
		Description:   rpcResp.Description,
		KeyFeatures:   rpcResp.KeyFeatures,
		ScientificName: rpcResp.ScientificName,
	}

	return resp, nil
}
