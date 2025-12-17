package hps

import (
	"encoding/json"
	"explorapal/app/model"

	"gorm.io/gorm"
)

// Achievement 成果记录表
type Achievement struct {
	model.BaseModel

	AchievementID int64  `gorm:"column:achievement_id;uniqueIndex;not null;comment:成果ID"`
	ProjectID     int64  `gorm:"column:project_id;index;not null;comment:项目ID"`
	UserID        int64  `gorm:"column:user_id;index;not null;comment:用户ID"`

	// 成果信息
	Type        string `gorm:"column:type;size:20;not null;comment:成果类型：report,documentary,poster"`
	Title       string `gorm:"column:title;size:200;not null;comment:成果标题"`
	Description string `gorm:"column:description;type:text;comment:成果描述"`

	// 内容数据
	Content     string `gorm:"column:content;type:longtext;not null;comment:成果内容JSON"`
	URL         string `gorm:"column:url;size:500;comment:成果URL"`
	FileSize    int64  `gorm:"column:file_size;comment:文件大小(字节)"`

	// 生成参数
	Style       string `gorm:"column:style;size:50;comment:风格参数"`
	Layout      string `gorm:"column:layout;size:50;comment:布局参数"`
	Length      string `gorm:"column:length;size:20;comment:时长参数"`

	// 状态
	Status      string `gorm:"column:status;size:20;default:completed;comment:状态：generating,completed,failed"`
	ErrorMsg    string `gorm:"column:error_msg;type:text;comment:错误信息"`

	// 统计信息
	ViewCount   int64 `gorm:"column:view_count;default:0;comment:查看次数"`
	LikeCount   int64 `gorm:"column:like_count;default:0;comment:点赞次数"`
	ShareCount  int64 `gorm:"column:share_count;default:0;comment:分享次数"`
}

// TableName 设置表名
func (Achievement) TableName() string {
	return "achievements"
}

// ResearchReport 研究报告结构
type ResearchReport struct {
	Title         string     `json:"title"`
	Abstract      string     `json:"abstract"`
	Introduction  string     `json:"introduction"`
	Methodology   string     `json:"methodology"`
	Findings      []Finding  `json:"findings"`
	Discussion    string     `json:"discussion"`
	Conclusion    string     `json:"conclusion"`
	References    []Reference `json:"references"`
	Visuals       []ReportVisual `json:"visuals"`
	ChildInsights string     `json:"child_insights"`
	NextSteps     []string   `json:"next_steps"`
}

// Finding 发现结构
type Finding struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Evidence    []string `json:"evidence"`
	Significance string  `json:"significance"`
}

// Reference 参考资料结构
type Reference struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	URL     string `json:"url,omitempty"`
	Credit  string `json:"credit"`
}

// ReportVisual 报告视觉元素结构
type ReportVisual struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

// DocumentaryScript 纪录片脚本结构
type DocumentaryScript struct {
	Title       string  `json:"title"`
	Duration    int32   `json:"duration"`
	Style       string  `json:"style"`
	Scenes      []Scene `json:"scenes"`
	Narration   string  `json:"narration"`
	Music       string  `json:"music"`
	Effects     []string `json:"effects"`
}

// Scene 场景结构
type Scene struct {
	SceneNumber int32    `json:"scene_number"`
	Duration    int32    `json:"duration"`
	Description string   `json:"description"`
	Visuals     []string `json:"visuals"`
	Narration   string   `json:"narration"`
	Transitions string   `json:"transitions,omitempty"`
}

// PosterDesign 海报设计结构
type PosterDesign struct {
	Title         string          `json:"title"`
	Style         string          `json:"style"`
	Layout        string          `json:"layout"`
	Sections      []PosterSection `json:"sections"`
	ColorScheme   ColorScheme    `json:"color_scheme"`
	Typography    Typography     `json:"typography"`
	VisualElements []VisualElement `json:"visual_elements"`
}

// PosterSection 海报区域结构
type PosterSection struct {
	Type    string  `json:"type"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Width   float64 `json:"width"`
	Height  float64 `json:"height"`
	Content string  `json:"content"`
	Style   string  `json:"style"`
}

// ColorScheme 配色方案结构
type ColorScheme struct {
	Primary   string   `json:"primary"`
	Secondary string   `json:"secondary"`
	Accent    string   `json:"accent"`
	Palette   []string `json:"palette"`
}

// Typography 字体设计结构
type Typography struct {
	TitleFont   string `json:"title_font"`
	BodyFont    string `json:"body_font"`
	HeadingFont string `json:"heading_font"`
	TitleSize   int32  `json:"title_size"`
	BodySize    int32  `json:"body_size"`
	HeadingSize int32  `json:"heading_size"`
}

// GetResearchReport 解析研究报告JSON
func (a *Achievement) GetResearchReport() (*ResearchReport, error) {
	if a.Content == "" {
		return &ResearchReport{}, nil
	}
	var report ResearchReport
	err := json.Unmarshal([]byte(a.Content), &report)
	return &report, err
}

// SetResearchReport 设置研究报告JSON
func (a *Achievement) SetResearchReport(report *ResearchReport) error {
	data, err := json.Marshal(report)
	if err != nil {
		return err
	}
	a.Content = string(data)
	return nil
}

// GetDocumentaryScript 解析纪录片脚本JSON
func (a *Achievement) GetDocumentaryScript() (*DocumentaryScript, error) {
	if a.Content == "" {
		return &DocumentaryScript{}, nil
	}
	var script DocumentaryScript
	err := json.Unmarshal([]byte(a.Content), &script)
	return &script, err
}

// SetDocumentaryScript 设置纪录片脚本JSON
func (a *Achievement) SetDocumentaryScript(script *DocumentaryScript) error {
	data, err := json.Marshal(script)
	if err != nil {
		return err
	}
	a.Content = string(data)
	return nil
}

// GetPosterDesign 解析海报设计JSON
func (a *Achievement) GetPosterDesign() (*PosterDesign, error) {
	if a.Content == "" {
		return &PosterDesign{}, nil
	}
	var design PosterDesign
	err := json.Unmarshal([]byte(a.Content), &design)
	return &design, err
}

// SetPosterDesign 设置海报设计JSON
func (a *Achievement) SetPosterDesign(design *PosterDesign) error {
	data, err := json.Marshal(design)
	if err != nil {
		return err
	}
	a.Content = string(data)
	return nil
}
