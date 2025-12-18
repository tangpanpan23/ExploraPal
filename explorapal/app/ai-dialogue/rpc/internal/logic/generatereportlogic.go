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
		l.Logger.Errorf("生成报告失败，使用默认报告: %v", err)
		// 返回默认的报告结果
		return l.getDefaultReportResponse(in.ProjectData, in.Category), nil
	}

	return &aidialogue.GenerateReportResp{
		Status:     200,
		Msg:        "生成报告成功",
		Title:      sanitizeUTF8(result.Title),
		Content:    sanitizeUTF8(result.Content),
		Abstract:   sanitizeUTF8(result.Abstract),
		Conclusion: sanitizeUTF8(result.Conclusion),
		NextSteps:  sanitizeUTF8Slice(result.NextSteps),
	}, nil
}

// getDefaultReportResponse 返回默认的报告生成结果
func (l *GenerateReportLogic) getDefaultReportResponse(projectData, category string) *aidialogue.GenerateReportResp {
	title := "探索报告"
	content := "这是基于你的探索经历生成的报告。由于AI服务暂时不可用，这里提供一个基本的报告模板。"
	abstract := "探索报告总结了孩子在学习过程中的观察、思考和发现。"
	conclusion := "通过这次探索，我们学到了很多新知识，希望继续保持好奇心和探索精神。"
	nextSteps := []string{
		"继续观察和记录",
		"提出更多问题",
		"尝试新的探索活动",
	}

	// 根据类别调整内容
	switch category {
	case "dinosaur":
		title = "恐龙探索报告"
		content = "在这次恐龙探索中，我们学习了恐龙的特征、生活习性和进化历程。"
		abstract = "通过观察恐龙化石和相关资料，我们对古代生物有了更深入的了解。"
		conclusion = "恐龙是地球历史上最神奇的生物之一，它们教会我们珍惜现在的生活。"
	case "rocket":
		title = "火箭探索报告"
		content = "在这次火箭探索中，我们学习了物理学原理和工程设计。"
		abstract = "通过火箭模型的制作和发射实验，我们了解了太空探索的基本原理。"
		conclusion = "太空探索需要勇气和智慧，让我们一起仰望星空，勇于探索。"
	}

	return &aidialogue.GenerateReportResp{
		Status:     200,
		Msg:        "生成报告成功（使用模拟响应）",
		Title:      sanitizeUTF8(title),
		Content:    sanitizeUTF8(content),
		Abstract:   sanitizeUTF8(abstract),
		Conclusion: sanitizeUTF8(conclusion),
		NextSteps:  sanitizeUTF8Slice(nextSteps),
	}
}

