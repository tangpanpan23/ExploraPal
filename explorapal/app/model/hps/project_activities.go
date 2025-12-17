package hps

import (
	"explorapal/app/model"

	"gorm.io/gorm"
)

// ProjectActivity 项目活动记录表
type ProjectActivity struct {
	model.BaseModel

	ActivityID int64  `gorm:"column:activity_id;uniqueIndex;not null;comment:活动ID"`
	ProjectID  int64  `gorm:"column:project_id;index;not null;comment:项目ID"`
	UserID     int64  `gorm:"column:user_id;index;not null;comment:用户ID"`

	// 活动信息
	Type        string `gorm:"column:type;size:50;not null;comment:活动类型：create_project,upload_image,recognize_image,generate_questions,select_question,speech_to_text,polish_note,generate_report,generate_documentary,generate_poster"`
	Description string `gorm:"column:description;type:text;not null;comment:活动描述"`

	// 元数据
	Metadata    string `gorm:"column:metadata;type:text;comment:活动元数据JSON"`
	IPAddress   string `gorm:"column:ip_address;size:45;comment:IP地址"`
	UserAgent   string `gorm:"column:user_agent;size:500;comment:用户代理"`
}

// TableName 设置表名
func (ProjectActivity) TableName() string {
	return "project_activities"
}

// ActivityType 活动类型常量
const (
	ActivityTypeCreateProject     = "create_project"
	ActivityTypeUploadImage       = "upload_image"
	ActivityTypeRecognizeImage    = "recognize_image"
	ActivityTypeGenerateQuestions = "generate_questions"
	ActivityTypeSelectQuestion    = "select_question"
	ActivityTypeSpeechToText      = "speech_to_text"
	ActivityTypePolishNote        = "polish_note"
	ActivityTypeGenerateReport    = "generate_report"
	ActivityTypeGenerateDocumentary = "generate_documentary"
	ActivityTypeGeneratePoster    = "generate_poster"
)
