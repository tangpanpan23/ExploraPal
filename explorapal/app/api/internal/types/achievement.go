package types

// 成果生成相关类型定义

type (
	GenerateReportReq struct {
		ProjectData string `json:"project_data" desc:"项目数据"`
		Category    string `json:"category" desc:"项目类别"`
	}

	GenerateReportResp struct {
		Title       string   `json:"title" desc:"报告标题"`
		Content     string   `json:"content" desc:"报告内容"`
		Abstract    string   `json:"abstract" desc:"摘要"`
		Conclusion  string   `json:"conclusion" desc:"结论"`
		NextSteps   []string `json:"next_steps" desc:"后续步骤建议"`
	}
)
