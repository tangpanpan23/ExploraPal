package types

// 提问引导相关类型定义

type (
	GenerateQuestionsReq struct {
		ContextInfo string `json:"context_info" desc:"上下文信息"`
		Category    string `json:"category" desc:"探索类别"`
		UserAge     int64  `json:"user_age,optional" desc:"用户年龄"`
	}

	GenerateQuestionsResp struct {
		Questions []QuestionItem `json:"questions" desc:"生成的问题列表"`
	}

	QuestionItem struct {
		Content    string `json:"content" desc:"问题内容"`
		Type       string `json:"type" desc:"问题类型：observation观察, reasoning推理, experiment实验, comparison比较"`
		Difficulty string `json:"difficulty" desc:"难度：basic基本, intermediate中级, advanced高级"`
		Purpose    string `json:"purpose" desc:"问题目的说明"`
	}
)
