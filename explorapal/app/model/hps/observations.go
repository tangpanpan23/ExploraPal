package hps

import (
	"encoding/json"
	"explorapal/app/model"
)

// Observation 观察记录表
type Observation struct {
	model.BaseModel

	ObservationID int64  `gorm:"column:observation_id;uniqueIndex;not null;comment:观察记录ID"`
	ProjectID     int64  `gorm:"column:project_id;index;not null;comment:项目ID"`
	UserID        int64  `gorm:"column:user_id;index;not null;comment:用户ID"`
	ImageURL      string `gorm:"column:image_url;size:5242880;not null;comment:图片URL或Base64数据"`
	ImageName     string `gorm:"column:image_name;size:200;comment:图片名称"`
	ImageType     string `gorm:"column:image_type;size:10;comment:图片类型：jpeg,png,jpg"`
	ImageSize     int64  `gorm:"column:image_size;comment:图片大小(字节)"`

	// 识别结果
	ObjectName     string  `gorm:"column:object_name;size:100;comment:识别对象名称"`
	Category       string  `gorm:"column:category;size:50;comment:类别"`
	Confidence     float64 `gorm:"column:confidence;comment:置信度"`
	Description    string  `gorm:"column:description;type:text;comment:描述"`
	KeyFeatures    string  `gorm:"column:key_features;type:text;comment:关键特征JSON数组"`
	ScientificName string  `gorm:"column:scientific_name;size:100;comment:学名"`

	// AI增强信息
	ARInfo         string `gorm:"column:ar_info;type:text;comment:AR信息JSON"`
	Suggestions    string `gorm:"column:suggestions;type:text;comment:AI建议JSON数组"`
	InterestingFacts string `gorm:"column:interesting_facts;type:text;comment:有趣事实JSON数组"`
}

// TableName 设置表名
func (Observation) TableName() string {
	return "observations"
}

// GetKeyFeatures 解析关键特征JSON
func (o *Observation) GetKeyFeatures() ([]string, error) {
	if o.KeyFeatures == "" {
		return []string{}, nil
	}
	var features []string
	err := json.Unmarshal([]byte(o.KeyFeatures), &features)
	return features, err
}

// SetKeyFeatures 设置关键特征JSON
func (o *Observation) SetKeyFeatures(features []string) error {
	data, err := json.Marshal(features)
	if err != nil {
		return err
	}
	o.KeyFeatures = string(data)
	return nil
}

// GetSuggestions 解析AI建议JSON
func (o *Observation) GetSuggestions() ([]string, error) {
	if o.Suggestions == "" {
		return []string{}, nil
	}
	var suggestions []string
	err := json.Unmarshal([]byte(o.Suggestions), &suggestions)
	return suggestions, err
}

// SetSuggestions 设置AI建议JSON
func (o *Observation) SetSuggestions(suggestions []string) error {
	data, err := json.Marshal(suggestions)
	if err != nil {
		return err
	}
	o.Suggestions = string(data)
	return nil
}

// GetInterestingFacts 解析有趣事实JSON
func (o *Observation) GetInterestingFacts() ([]string, error) {
	if o.InterestingFacts == "" {
		return []string{}, nil
	}
	var facts []string
	err := json.Unmarshal([]byte(o.InterestingFacts), &facts)
	return facts, err
}

// SetInterestingFacts 设置有趣事实JSON
func (o *Observation) SetInterestingFacts(facts []string) error {
	data, err := json.Marshal(facts)
	if err != nil {
		return err
	}
	o.InterestingFacts = string(data)
	return nil
}

// ARInformation AR信息结构
type ARInformation struct {
	Hotspots []ARHotspot `json:"hotspots"`
	Labels   []ARLabel   `json:"labels"`
}

// ARHotspot AR热点
type ARHotspot struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Type    string  `json:"type"`
}

// ARLabel AR标签
type ARLabel struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Text  string  `json:"text"`
	Color string  `json:"color"`
}

// GetARInfo 解析AR信息JSON
func (o *Observation) GetARInfo() (*ARInformation, error) {
	if o.ARInfo == "" {
		return &ARInformation{}, nil
	}
	var arInfo ARInformation
	err := json.Unmarshal([]byte(o.ARInfo), &arInfo)
	return &arInfo, err
}

// SetARInfo 设置AR信息JSON
func (o *Observation) SetARInfo(arInfo *ARInformation) error {
	data, err := json.Marshal(arInfo)
	if err != nil {
		return err
	}
	o.ARInfo = string(data)
	return nil
}
