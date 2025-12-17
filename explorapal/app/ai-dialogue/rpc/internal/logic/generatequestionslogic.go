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
		l.Logger.Errorf("生成问题失败: %v", err)
		return &aidialogue.GenerateQuestionsResp{
			Status: 500,
			Msg:    "生成问题失败",
		}, err
	}

	var questionList []*aidialogue.Question
	for _, q := range questions {
		questionList = append(questionList, &aidialogue.Question{
			Content:    q.Content,
			Type:       q.Type,
			Difficulty: q.Difficulty,
			Purpose:    q.Purpose,
		})
	}

	return &aidialogue.GenerateQuestionsResp{
		Status:    200,
		Msg:       "生成问题成功",
		Questions: questionList,
	}, nil
}
