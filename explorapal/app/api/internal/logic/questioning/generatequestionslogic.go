package questioning

import (
	"context"
	"time"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateQuestionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成引导问题
func NewGenerateQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateQuestionsLogic {
	return &GenerateQuestionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateQuestionsLogic) GenerateQuestions(req *types.GenerateQuestionsReq) (resp *types.GenerateQuestionsResp, err error) {
	// 调用AI服务生成问题
	aiQuestions, err := l.svcCtx.AIClient.GenerateQuestions(l.ctx, req.ContextInfo, req.Category)
	if err != nil {
		l.Errorf("AI生成问题失败: %v", err)
		return nil, err
	}

	// 转换AI返回的问题为API格式
	questions := make([]types.Question, len(aiQuestions))
	for i, aiQuestion := range aiQuestions {
		questions[i] = types.Question{
			QuestionId:       time.Now().UnixNano()/1000000 + int64(i), // 生成问题ID
			Content:          aiQuestion.Content,
			Type:             aiQuestion.Type,
			Difficulty:       aiQuestion.Difficulty,
			Purpose:          aiQuestion.Purpose,
			Hints:            []string{}, // 可以后续添加提示
			ExpectedThinking: "",        // 可以后续添加期望思考
		}
	}

	// 如果没有生成问题，提供默认问题
	if len(questions) == 0 {
		questions = l.getDefaultQuestions(req.Category)
	}

	resp = &types.GenerateQuestionsResp{
		Questions: questions,
	}

	l.Infof("成功生成 %d 个问题，项目ID: %d, 类别: %s", len(questions), req.ProjectId, req.Category)

	return resp, nil
}

// getDefaultQuestions 获取默认问题（当AI服务不可用时使用）
func (l *GenerateQuestionsLogic) getDefaultQuestions(category string) []types.Question {
	baseID := time.Now().UnixNano() / 1000000

	switch category {
	case "dinosaur":
		return []types.Question{
			{
				QuestionId:       baseID + 1,
				Content:          "恐龙的牙齿形状可以告诉我们什么信息？",
				Type:             "observation",
				Difficulty:       "basic",
				Purpose:          "培养观察力和推理能力",
				Hints:            []string{"仔细观察牙齿的形状和大小"},
				ExpectedThinking: "不同形状的牙齿对应不同的饮食习惯",
			},
			{
				QuestionId:       baseID + 2,
				Content:          "为什么有些恐龙有角或者骨板？",
				Type:             "reasoning",
				Difficulty:       "intermediate",
				Purpose:          "理解适者生存的进化原理",
				Hints:            []string{"想想这些特征有什么作用"},
				ExpectedThinking: "这些特征可能帮助恐龙生存和繁衍",
			},
		}
	case "rocket":
		return []types.Question{
			{
				QuestionId:       baseID + 1,
				Content:          "火箭为什么需要燃料？",
				Type:             "reasoning",
				Difficulty:       "basic",
				Purpose:          "理解牛顿第三定律",
				Hints:            []string{"想想作用力和反作用力"},
				ExpectedThinking: "燃料燃烧产生推力，推动火箭前进",
			},
		}
	default:
		return []types.Question{
			{
				QuestionId:       baseID + 1,
				Content:          "你观察到了什么有趣的地方？",
				Type:             "observation",
				Difficulty:       "basic",
				Purpose:          "培养观察力",
				Hints:            []string{"用眼睛仔细看"},
				ExpectedThinking: "观察是探索的第一步",
			},
		}
	}
}
