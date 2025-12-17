package hps

import (
	"encoding/json"
	"time"

	"explorapal/app/model"
)

// Question 问题记录表
type Question struct {
	model.BaseModel

	QuestionID   int64  `gorm:"column:question_id;uniqueIndex;not null;comment:问题ID"`
	ProjectID    int64  `gorm:"column:project_id;index;not null;comment:项目ID"`
	UserID       int64  `gorm:"column:user_id;index;not null;comment:用户ID"`
	ObservationID int64 `gorm:"column:observation_id;index;comment:关联的观察记录ID"`

	// 问题信息
	Content     string `gorm:"column:content;type:text;not null;comment:问题内容"`
	Type        string `gorm:"column:type;size:20;not null;comment:问题类型：observation,reasoning,experiment,comparison"`
	Difficulty  string `gorm:"column:difficulty;size:20;default:basic;comment:难度级别：basic,intermediate,advanced"`
	Purpose     string `gorm:"column:purpose;type:text;comment:问题目的"`
	Hints       string `gorm:"column:hints;type:text;comment:提示JSON数组"`
	ExpectedThinking string `gorm:"column:expected_thinking;type:text;comment:期望的思考方向"`

	// AI回答
	AIAnswer         string `gorm:"column:ai_answer;type:text;comment:AI回答"`
	KeyPoints        string `gorm:"column:key_points;type:text;comment:关键要点JSON数组"`
	Examples         string `gorm:"column:examples;type:text;comment:举例说明JSON数组"`
	Analogies        string `gorm:"column:analogies;type:text;comment:类比JSON数组"`
	VisualAids       string `gorm:"column:visual_aids;type:text;comment:视觉辅助建议JSON数组"`

	// 后续问题和活动
	FollowUpQuestions string `gorm:"column:follow_up_questions;type:text;comment:后续问题建议JSON数组"`
	ThinkingPrompts   string `gorm:"column:thinking_prompts;type:text;comment:思考提示JSON数组"`
	Activities        string `gorm:"column:activities;type:text;comment:建议活动JSON数组"`

	// 用户响应
	UserResponse string `gorm:"column:user_response;type:text;comment:用户回答"`
	ResponseTime *time.Time `gorm:"column:response_time;comment:回答时间"`
}

// TableName 设置表名
func (Question) TableName() string {
	return "questions"
}

// GetHints 解析提示JSON
func (q *Question) GetHints() ([]string, error) {
	if q.Hints == "" {
		return []string{}, nil
	}
	var hints []string
	err := json.Unmarshal([]byte(q.Hints), &hints)
	return hints, err
}

// SetHints 设置提示JSON
func (q *Question) SetHints(hints []string) error {
	data, err := json.Marshal(hints)
	if err != nil {
		return err
	}
	q.Hints = string(data)
	return nil
}

// GetKeyPoints 解析关键要点JSON
func (q *Question) GetKeyPoints() ([]string, error) {
	if q.KeyPoints == "" {
		return []string{}, nil
	}
	var points []string
	err := json.Unmarshal([]byte(q.KeyPoints), &points)
	return points, err
}

// SetKeyPoints 设置关键要点JSON
func (q *Question) SetKeyPoints(points []string) error {
	data, err := json.Marshal(points)
	if err != nil {
		return err
	}
	q.KeyPoints = string(data)
	return nil
}

// GetExamples 解析举例说明JSON
func (q *Question) GetExamples() ([]string, error) {
	if q.Examples == "" {
		return []string{}, nil
	}
	var examples []string
	err := json.Unmarshal([]byte(q.Examples), &examples)
	return examples, err
}

// SetExamples 设置举例说明JSON
func (q *Question) SetExamples(examples []string) error {
	data, err := json.Marshal(examples)
	if err != nil {
		return err
	}
	q.Examples = string(data)
	return nil
}

// GetAnalogies 解析类比JSON
func (q *Question) GetAnalogies() ([]string, error) {
	if q.Analogies == "" {
		return []string{}, nil
	}
	var analogies []string
	err := json.Unmarshal([]byte(q.Analogies), &analogies)
	return analogies, err
}

// SetAnalogies 设置类比JSON
func (q *Question) SetAnalogies(analogies []string) error {
	data, err := json.Marshal(analogies)
	if err != nil {
		return err
	}
	q.Analogies = string(data)
	return nil
}

// GetVisualAids 解析视觉辅助建议JSON
func (q *Question) GetVisualAids() ([]string, error) {
	if q.VisualAids == "" {
		return []string{}, nil
	}
	var aids []string
	err := json.Unmarshal([]byte(q.VisualAids), &aids)
	return aids, err
}

// SetVisualAids 设置视觉辅助建议JSON
func (q *Question) SetVisualAids(aids []string) error {
	data, err := json.Marshal(aids)
	if err != nil {
		return err
	}
	q.VisualAids = string(data)
	return nil
}

// GetFollowUpQuestions 解析后续问题建议JSON
func (q *Question) GetFollowUpQuestions() ([]string, error) {
	if q.FollowUpQuestions == "" {
		return []string{}, nil
	}
	var questions []string
	err := json.Unmarshal([]byte(q.FollowUpQuestions), &questions)
	return questions, err
}

// SetFollowUpQuestions 设置后续问题建议JSON
func (q *Question) SetFollowUpQuestions(questions []string) error {
	data, err := json.Marshal(questions)
	if err != nil {
		return err
	}
	q.FollowUpQuestions = string(data)
	return nil
}

// GetThinkingPrompts 解析思考提示JSON
func (q *Question) GetThinkingPrompts() ([]string, error) {
	if q.ThinkingPrompts == "" {
		return []string{}, nil
	}
	var prompts []string
	err := json.Unmarshal([]byte(q.ThinkingPrompts), &prompts)
	return prompts, err
}

// SetThinkingPrompts 设置思考提示JSON
func (q *Question) SetThinkingPrompts(prompts []string) error {
	data, err := json.Marshal(prompts)
	if err != nil {
		return err
	}
	q.ThinkingPrompts = string(data)
	return nil
}

// Activity 活动结构
type Activity struct {
	Type        string   `json:"type"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Materials   []string `json:"materials"`
	Steps       []string `json:"steps"`
	Duration    int32    `json:"duration"`
	Difficulty  string   `json:"difficulty"`
}

// GetActivities 解析建议活动JSON
func (q *Question) GetActivities() ([]Activity, error) {
	if q.Activities == "" {
		return []Activity{}, nil
	}
	var activities []Activity
	err := json.Unmarshal([]byte(q.Activities), &activities)
	return activities, err
}

// SetActivities 设置建议活动JSON
func (q *Question) SetActivities(activities []Activity) error {
	data, err := json.Marshal(activities)
	if err != nil {
		return err
	}
	q.Activities = string(data)
	return nil
}
