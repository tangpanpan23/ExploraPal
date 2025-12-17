package hps

import (
	"encoding/json"
	"explorapal/app/model"

	"gorm.io/gorm"
)

// Project 探索项目表
type Project struct {
	model.BaseModel

	ProjectID     int64    `gorm:"column:project_id;uniqueIndex;not null;comment:项目ID"`
	ProjectCode   string   `gorm:"column:project_code;size:50;uniqueIndex;comment:项目编码"`
	UserID        int64    `gorm:"column:user_id;index;not null;comment:用户ID"`
	Title         string   `gorm:"column:title;size:200;not null;comment:项目标题"`
	Description   string   `gorm:"column:description;type:text;comment:项目描述"`
	Category      string   `gorm:"column:category;size:50;not null;comment:项目类别"`
	Status        string   `gorm:"column:status;size:20;default:active;comment:状态：active,completed,paused"`
	Progress      int32    `gorm:"column:progress;default:0;comment:进度百分比(0-100)"`
	Tags          string   `gorm:"column:tags;type:text;comment:标签JSON数组"`
	LastActivityAt *time.Time `gorm:"column:last_activity_at;comment:最后活动时间"`
}

// TableName 设置表名
func (Project) TableName() string {
	return "projects"
}

// GetTags 解析标签JSON
func (p *Project) GetTags() ([]string, error) {
	if p.Tags == "" {
		return []string{}, nil
	}
	var tags []string
	err := json.Unmarshal([]byte(p.Tags), &tags)
	return tags, err
}

// SetTags 设置标签JSON
func (p *Project) SetTags(tags []string) error {
	data, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	p.Tags = string(data)
	return nil
}

// BeforeCreate 创建前的钩子
func (p *Project) BeforeCreate(tx *gorm.DB) error {
	if p.Status == "" {
		p.Status = "active"
	}
	if p.Progress == 0 {
		p.Progress = 0
	}
	return nil
}
