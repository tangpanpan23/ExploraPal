package logic

import (
	"context"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/ai-dialogue/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GenerateReportLogic) GenerateReport(in *aidialogue.GenerateReportReq) (*aidialogue.GenerateReportResp, error) {
	// TODO: 实现报告生成逻辑
	result, err := l.svcCtx.AIClient.GenerateReport(l.ctx, in.ProjectData)
	if err != nil {
		l.Logger.Errorf("生成报告失败: %v", err)
		return &aidialogue.GenerateReportResp{
			Status: 500,
			Msg:    "生成报告失败",
		}, err
	}

	return &aidialogue.GenerateReportResp{
		Status:     200,
		Msg:        "生成报告成功",
		Title:      result.Title,
		Content:    result.Content,
		Abstract:   result.Abstract,
		Conclusion: result.Conclusion,
		NextSteps:  result.NextSteps,
	}, nil
}
}
