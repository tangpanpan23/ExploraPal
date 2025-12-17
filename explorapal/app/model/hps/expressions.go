package hps

import (
	"encoding/json"
	"explorapal/app/model"
)

// Expression 表达记录表
type Expression struct {
	model.BaseModel

	ExpressionID int64  `gorm:"column:expression_id;uniqueIndex;not null;comment:表达记录ID"`
	ProjectID    int64  `gorm:"column:project_id;index;not null;comment:项目ID"`
	UserID       int64  `gorm:"column:user_id;index;not null;comment:用户ID"`
	QuestionID   int64  `gorm:"column:question_id;index;comment:关联的问题ID"`

	// 内容信息
	Type        string  `gorm:"column:type;size:20;not null;comment:类型：speech,text,note"`
	RawContent  string  `gorm:"column:raw_content;type:text;not null;comment:原始内容"`
	Language    string  `gorm:"column:language;size:10;default:zh-CN;comment:语言代码"`

	// 语音转文字相关
	AudioURL    string  `gorm:"column:audio_url;size:500;comment:音频URL"`
	AudioFormat string  `gorm:"column:audio_format;size:10;comment:音频格式：wav,mp3,m4a"`
	AudioSize   int64   `gorm:"column:audio_size;comment:音频大小(字节)"`
	Duration    float64 `gorm:"column:duration;comment:音频时长(秒)"`
	Confidence  float64 `gorm:"column:confidence;comment:识别置信度"`

	// AI润色结果
	PolishedTitle       string `gorm:"column:polished_title;size:200;comment:润色后的标题"`
	PolishedSummary     string `gorm:"column:polished_summary;type:text;comment:内容总结"`
	PolishedKeyPoints   string `gorm:"column:polished_key_points;type:text;comment:关键要点JSON数组"`
	PolishedConcepts    string `gorm:"column:polished_concepts;type:text;comment:科学概念JSON数组"`
	PolishedQuestions   string `gorm:"column:polished_questions;type:text;comment:引发的疑问JSON数组"`
	PolishedConnections string `gorm:"column:polished_connections;type:text;comment:关联知识JSON数组"`
	PolishedVisuals     string `gorm:"column:polished_visuals;type:text;comment:视觉元素JSON数组"`
	PolishedFormatted   string `gorm:"column:polished_formatted;type:text;comment:格式化文本"`

	// AI建议
	Suggestions  string `gorm:"column:suggestions;type:text;comment:改进建议JSON数组"`
	KeyLearnings string `gorm:"column:key_learnings;type:text;comment:关键学习点JSON数组"`
}

// TableName 设置表名
func (Expression) TableName() string {
	return "expressions"
}

// GetPolishedKeyPoints 解析润色后的关键要点JSON
func (e *Expression) GetPolishedKeyPoints() ([]string, error) {
	if e.PolishedKeyPoints == "" {
		return []string{}, nil
	}
	var points []string
	err := json.Unmarshal([]byte(e.PolishedKeyPoints), &points)
	return points, err
}

// SetPolishedKeyPoints 设置润色后的关键要点JSON
func (e *Expression) SetPolishedKeyPoints(points []string) error {
	data, err := json.Marshal(points)
	if err != nil {
		return err
	}
	e.PolishedKeyPoints = string(data)
	return nil
}

// GetPolishedConcepts 解析润色后的科学概念JSON
func (e *Expression) GetPolishedConcepts() ([]string, error) {
	if e.PolishedConcepts == "" {
		return []string{}, nil
	}
	var concepts []string
	err := json.Unmarshal([]byte(e.PolishedConcepts), &concepts)
	return concepts, err
}

// SetPolishedConcepts 设置润色后的科学概念JSON
func (e *Expression) SetPolishedConcepts(concepts []string) error {
	data, err := json.Marshal(concepts)
	if err != nil {
		return err
	}
	e.PolishedConcepts = string(data)
	return nil
}

// GetPolishedQuestions 解析润色后的疑问JSON
func (e *Expression) GetPolishedQuestions() ([]string, error) {
	if e.PolishedQuestions == "" {
		return []string{}, nil
	}
	var questions []string
	err := json.Unmarshal([]byte(e.PolishedQuestions), &questions)
	return questions, err
}

// SetPolishedQuestions 设置润色后的疑问JSON
func (e *Expression) SetPolishedQuestions(questions []string) error {
	data, err := json.Marshal(questions)
	if err != nil {
		return err
	}
	e.PolishedQuestions = string(data)
	return nil
}

// GetPolishedConnections 解析润色后的关联知识JSON
func (e *Expression) GetPolishedConnections() ([]string, error) {
	if e.PolishedConnections == "" {
		return []string{}, nil
	}
	var connections []string
	err := json.Unmarshal([]byte(e.PolishedConnections), &connections)
	return connections, err
}

// SetPolishedConnections 设置润色后的关联知识JSON
func (e *Expression) SetPolishedConnections(connections []string) error {
	data, err := json.Marshal(connections)
	if err != nil {
		return err
	}
	e.PolishedConnections = string(data)
	return nil
}

// VisualElement 视觉元素结构
type VisualElement struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Data        string `json:"data"`
	Position    string `json:"position"`
}

// GetPolishedVisuals 解析润色后的视觉元素JSON
func (e *Expression) GetPolishedVisuals() ([]VisualElement, error) {
	if e.PolishedVisuals == "" {
		return []VisualElement{}, nil
	}
	var visuals []VisualElement
	err := json.Unmarshal([]byte(e.PolishedVisuals), &visuals)
	return visuals, err
}

// SetPolishedVisuals 设置润色后的视觉元素JSON
func (e *Expression) SetPolishedVisuals(visuals []VisualElement) error {
	data, err := json.Marshal(visuals)
	if err != nil {
		return err
	}
	e.PolishedVisuals = string(data)
	return nil
}

// GetSuggestions 解析改进建议JSON
func (e *Expression) GetSuggestions() ([]string, error) {
	if e.Suggestions == "" {
		return []string{}, nil
	}
	var suggestions []string
	err := json.Unmarshal([]byte(e.Suggestions), &suggestions)
	return suggestions, err
}

// SetSuggestions 设置改进建议JSON
func (e *Expression) SetSuggestions(suggestions []string) error {
	data, err := json.Marshal(suggestions)
	if err != nil {
		return err
	}
	e.Suggestions = string(data)
	return nil
}

// GetKeyLearnings 解析关键学习点JSON
func (e *Expression) GetKeyLearnings() ([]string, error) {
	if e.KeyLearnings == "" {
		return []string{}, nil
	}
	var learnings []string
	err := json.Unmarshal([]byte(e.KeyLearnings), &learnings)
	return learnings, err
}

// SetKeyLearnings 设置关键学习点JSON
func (e *Expression) SetKeyLearnings(learnings []string) error {
	data, err := json.Marshal(learnings)
	if err != nil {
		return err
	}
	e.KeyLearnings = string(data)
	return nil
}
