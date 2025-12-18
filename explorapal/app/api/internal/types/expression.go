package types

// 表达阶段相关类型定义

type (
	PolishNoteReq struct {
		RawContent  string `json:"raw_content" desc:"原始笔记内容"`
		ContextInfo string `json:"context_info" desc:"上下文信息"`
		Category    string `json:"category" desc:"探索类别"`
		UserAge     int64  `json:"user_age,optional" desc:"用户年龄"`
	}

	PolishNoteResp struct {
		Title         string   `json:"title" desc:"润色后的标题"`
		Summary       string   `json:"summary" desc:"内容摘要"`
		KeyPoints     []string `json:"key_points" desc:"关键要点"`
		FormattedText string   `json:"formatted_text" desc:"格式化文本"`
		Suggestions   []string `json:"suggestions" desc:"改进建议"`
	}
)
