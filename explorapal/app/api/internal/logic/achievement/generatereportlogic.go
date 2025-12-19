package achievement

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成研究简报
func NewGenerateReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateReportLogic {
	return &GenerateReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateReportLogic) GenerateReport(req *types.GenerateReportReq) (resp *types.GenerateReportResp, err error) {
	// 调用AI服务生成研究报告
	aiReport, err := l.svcCtx.AIClient.GenerateReport(l.ctx, req.ProjectData)
	if err != nil {
		l.Errorf("AI生成报告失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp = &types.GenerateReportResp{
		Title:       aiReport.Title,
		Content:     aiReport.Content,
		Abstract:    aiReport.Abstract,
		Conclusion:  aiReport.Conclusion,
		NextSteps:   aiReport.NextSteps,
	}

	l.Infof("研究报告生成完成: 标题=%s", aiReport.Title)

	return resp, nil
}
