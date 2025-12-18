package logic

import (
	"context"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type GenerateReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateReportLogic {
	return &GenerateReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateReportLogic) GenerateReport(req *types.GenerateReportReq) (resp *types.GenerateReportResp, err error) {
	// 调用AI对话RPC服务生成报告
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9002"}, // AI对话RPC服务地址
		Timeout:   60000,                      // 60秒超时，与AI服务超时保持一致
	})
	if err != nil {
		l.Logger.Errorf("创建AI RPC客户端失败: %v", err)
		return nil, err
	}

	aiClient := aidialogue.NewAIDialogueServiceClient(client.Conn())

	rpcResp, err := aiClient.GenerateReport(l.ctx, &aidialogue.GenerateReportReq{
		ProjectData: req.ProjectData,
		Category:    req.Category,
	})
	if err != nil {
		l.Logger.Errorf("调用AI生成报告服务失败: %v", err)
		return nil, err
	}

	resp = &types.GenerateReportResp{
		Title:      rpcResp.Title,
		Content:    rpcResp.Content,
		Abstract:   rpcResp.Abstract,
		Conclusion: rpcResp.Conclusion,
		NextSteps:  rpcResp.NextSteps,
	}

	return resp, nil
}
