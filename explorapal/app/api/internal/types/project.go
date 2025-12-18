package types

// 项目管理相关类型定义

type (
	CreateProjectReq struct {
		UserId      int64    `json:"user_id" desc:"用户ID"`
		Title       string   `json:"title" desc:"项目标题"`
		Description string   `json:"description,optional" desc:"项目描述"`
		Category    string   `json:"category" desc:"项目类别：dinosaur,rocket,minecraft等"`
		Tags        []string `json:"tags,optional" desc:"标签"`
	}

	CreateProjectResp struct {
		ProjectId   int64  `json:"project_id" desc:"项目ID"`
		ProjectCode string `json:"project_code" desc:"项目编码"`
		Status      string `json:"status" desc:"项目状态"`
	}

	GetProjectListReq struct {
		UserId   int64  `json:"user_id" desc:"用户ID"`
		Category string `json:"category,optional" desc:"项目类别筛选"`
		Status   string `json:"status,optional" desc:"状态筛选：active,completed,paused"`
		PageSize int64  `json:"page_size,optional,default=10" desc:"每页条数"`
		Page     int64  `json:"page,optional,default=1" desc:"页码"`
	}

	GetProjectListResp struct {
		List     []ProjectInfo `json:"list" desc:"项目列表"`
		Total    int64         `json:"total" desc:"总数"`
		PageSize int64         `json:"page_size" desc:"每页条数"`
		Page     int64         `json:"page" desc:"页码"`
	}

	ProjectInfo struct {
		ProjectId    int64    `json:"project_id" desc:"项目ID"`
		ProjectCode  string   `json:"project_code" desc:"项目编码"`
		Title        string   `json:"title" desc:"项目标题"`
		Description  string   `json:"description" desc:"项目描述"`
		Category     string   `json:"category" desc:"项目类别"`
		Status       string   `json:"status" desc:"项目状态"`
		Progress     int32    `json:"progress" desc:"进度百分比"`
		CreateTime   string   `json:"create_time" desc:"创建时间"`
		UpdateTime   string   `json:"update_time" desc:"更新时间"`
		LastActivity string   `json:"last_activity" desc:"最后活动时间"`
		Tags         []string `json:"tags" desc:"标签"`
	}

	GetProjectDetailReq struct {
		ProjectId int64 `json:"project_id" desc:"项目ID"`
		UserId    int64 `json:"user_id" desc:"用户ID"`
	}

	GetProjectDetailResp struct {
		Project      ProjectDetail        `json:"project" desc:"项目详情"`
		Activities   []ProjectActivity    `json:"activities" desc:"项目活动记录"`
		Achievements []ProjectAchievement `json:"achievements" desc:"项目成果"`
	}

	ProjectDetail struct {
		ProjectId    int64             `json:"project_id" desc:"项目ID"`
		ProjectCode  string            `json:"project_code" desc:"项目编码"`
		Title        string            `json:"title" desc:"项目标题"`
		Description  string            `json:"description" desc:"项目描述"`
		Category     string            `json:"category" desc:"项目类别"`
		Status       string            `json:"status" desc:"项目状态"`
		Progress     int32             `json:"progress" desc:"进度百分比"`
		CreateTime   string            `json:"create_time" desc:"创建时间"`
		UpdateTime   string            `json:"update_time" desc:"更新时间"`
		LastActivity string            `json:"last_activity" desc:"最后活动时间"`
		Tags         []string          `json:"tags" desc:"标签"`
		Observations []ObservationInfo `json:"observations" desc:"观察记录"`
		Questions    []QuestionInfo    `json:"questions" desc:"提问记录"`
		Expressions  []ExpressionInfo  `json:"expressions" desc:"表达记录"`
	}

	ObservationInfo struct {
		ObservationId int64  `json:"observation_id" desc:"观察ID"`
		ImageUrl      string `json:"image_url" desc:"图片URL"`
		Recognition   string `json:"recognition" desc:"识别结果"`
		CreateTime    string `json:"create_time" desc:"创建时间"`
	}

	QuestionInfo struct {
		QuestionId   int64  `json:"question_id" desc:"问题ID"`
		Question     string `json:"question" desc:"问题内容"`
		Answer       string `json:"answer" desc:"AI回答"`
		UserResponse string `json:"user_response,optional" desc:"用户回答"`
		CreateTime   string `json:"create_time" desc:"创建时间"`
	}

	ExpressionInfo struct {
		ExpressionId int64  `json:"expression_id" desc:"表达ID"`
		Content      string `json:"content" desc:"内容"`
		Type         string `json:"type" desc:"类型：speech,text,note"`
		PolishedNote string `json:"polished_note,optional" desc:"润色后的笔记"`
		CreateTime   string `json:"create_time" desc:"创建时间"`
	}

	ProjectActivity struct {
		ActivityId  int64  `json:"activity_id" desc:"活动ID"`
		Type        string `json:"type" desc:"活动类型"`
		Description string `json:"description" desc:"活动描述"`
		CreateTime  string `json:"create_time" desc:"创建时间"`
	}

	ProjectAchievement struct {
		AchievementId int64  `json:"achievement_id" desc:"成果ID"`
		Type          string `json:"type" desc:"成果类型：report,documentary,poster"`
		Title         string `json:"title" desc:"成果标题"`
		Content       string `json:"content" desc:"成果内容"`
		Url           string `json:"url,optional" desc:"成果URL"`
		CreateTime    string `json:"create_time" desc:"创建时间"`
	}

	UpdateProjectStatusReq struct {
		ProjectId int64  `json:"project_id" desc:"项目ID"`
		UserId    int64  `json:"user_id" desc:"用户ID"`
		Status    string `json:"status" desc:"新状态：active,completed,paused"`
	}

	CommonStatusResp struct {
		Code    int32  `json:"code" desc:"响应码"`
		Message string `json:"message" desc:"响应消息"`
	}
)
