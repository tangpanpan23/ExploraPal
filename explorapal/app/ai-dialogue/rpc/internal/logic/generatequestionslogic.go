package logic

import (
	"context"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/ai-dialogue/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateQuestionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateQuestionsLogic {
	return &GenerateQuestionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateQuestionsLogic) GenerateQuestions(in *aidialogue.GenerateQuestionsReq) (*aidialogue.GenerateQuestionsResp, error) {
	// TODO: 实现问题生成逻辑
	questions, err := l.svcCtx.AIClient.GenerateQuestions(l.ctx, in.ContextInfo, in.Category)
	if err != nil {
		l.Logger.Errorf("生成问题失败，使用默认问题: %v", err)
		// 返回默认问题列表
		return l.getDefaultQuestionsResponse(in.Category), nil
	}

	var questionList []*aidialogue.Question
	for _, q := range questions {
		questionList = append(questionList, &aidialogue.Question{
			Content:    sanitizeUTF8(q.Content),
			Type:       sanitizeUTF8(q.Type),
			Difficulty: sanitizeUTF8(q.Difficulty),
			Purpose:    sanitizeUTF8(q.Purpose),
		})
	}

	return &aidialogue.GenerateQuestionsResp{
		Status:    200,
		Msg:       "生成问题成功",
		Questions: questionList,
	}, nil
}

// getDefaultQuestionsResponse 根据类别返回默认问题响应
func (l *GenerateQuestionsLogic) getDefaultQuestionsResponse(category string) *aidialogue.GenerateQuestionsResp {
	var questions []*aidialogue.Question

	switch category {
	case "dinosaur":
		questions = []*aidialogue.Question{
			{
				Content:    "你觉得三角龙的角有什么用处？",
				Type:       "reasoning",
				Difficulty: "intermediate",
				Purpose:    "培养推理能力和想象力",
			},
			{
				Content:    "我们可以从化石上看到三角龙吃什么吗？",
				Type:       "observation",
				Difficulty: "basic",
				Purpose:    "学习观察细节",
			},
			{
				Content:    "如果我们能见到活的三角龙，你最想问它什么问题？",
				Type:       "experiment",
				Difficulty: "advanced",
				Purpose:    "激发科学探索精神",
			},
		}
	default:
		questions = []*aidialogue.Question{
			{
				Content:    "你看到了什么有趣的东西？",
				Type:       "observation",
				Difficulty: "basic",
				Purpose:    "激发观察兴趣",
			},
			{
				Content:    "为什么会这样呢？",
				Type:       "reasoning",
				Difficulty: "intermediate",
				Purpose:    "培养思考能力",
			},
			{
				Content:    "我们可以做个小实验验证吗？",
				Type:       "experiment",
				Difficulty: "advanced",
				Purpose:    "实践探索精神",
			},
		}
	}

	return &aidialogue.GenerateQuestionsResp{
		Status:    200,
		Msg:       "生成问题成功（使用模拟响应）",
		Questions: questions,
	}
}

