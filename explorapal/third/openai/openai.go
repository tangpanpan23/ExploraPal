package openai

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

// Client 内部AI服务客户端 (兼容OpenAI接口)
type Client struct {
	client *openai.Client
	config *Config
}

// Config 内部AI服务配置
type Config struct {
	TAL_MLOPS_APP_ID  string  `json:"talMLOpsAppId"`         // TAL MLOps应用ID
	TAL_MLOPS_APP_KEY string  `json:"talMLOpsAppKey"`        // TAL MLOps应用密钥
	BaseURL           string  `json:"baseURL,omitempty"`     // 内部AI服务端点，默认: http://ai-service.tal.com/openai-compatible/v1
	Timeout           int     `json:"timeout,omitempty"`     // 超时时间(秒)
	MaxTokens         int     `json:"maxTokens,omitempty"`   // 最大token数
	Temperature       float32 `json:"temperature,omitempty"` // 温度参数
}

// 内部AI服务模型映射 (通过TAL MLOps平台)
const (
	// 图像分析 - 使用多模态视觉理解模型
	ModelImageAnalysis = "qwen3-vl-plus" // 视觉理解，支持思考模式

	// 问题生成和笔记润色 - 使用快速Flash模型
	ModelTextGeneration = "qwen-flash" // 思考+非思考模式融合

	// 复杂推理和报告生成 - 使用Max模型
	ModelAdvancedReasoning = "qwen3-max" // 智能体编程和工具调用优化

	// 语音交互 - 使用Omni多模态模型
	ModelVoiceInteraction = "qwen3-omni-flash" // 多模态语音处理

	// 备用模型
	ModelImageAnalysisBackup     = "qwen3-vl-235b-a22b-instruct" // 备用的视觉模型
	ModelTextGenerationBackup    = "qwen-turbo"                  // 备用的快速模型
	ModelAdvancedReasoningBackup = "qwen-max"                    // 备用的推理模型
	ModelVoiceInteractionBackup  = "qwen3-omni-flash"            // 语音交互备用模型
)

// NewClient 创建内部AI服务客户端
func NewClient(config *Config) *Client {
	// 构建认证token
	token := fmt.Sprintf("%s:%s", config.TAL_MLOPS_APP_ID, config.TAL_MLOPS_APP_KEY)

	// 设置内部AI服务端点
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "http://ai-service.tal.com/openai-compatible/v1"
	}

	// 创建客户端配置
	clientConfig := openai.DefaultConfig(token)
	clientConfig.BaseURL = baseURL
	clientConfig.APIType = openai.APITypeOpenAI

	// 设置自定义HTTP headers用于认证
	clientConfig.HTTPClient.Transport = &customTransport{
		base:  clientConfig.HTTPClient.Transport,
		token: fmt.Sprintf("Bearer %s", token),
	}

	return &Client{
		client: openai.NewClientWithConfig(clientConfig),
		config: config,
	}
}

// customTransport 自定义传输层，添加Authorization header
type customTransport struct {
	base  http.RoundTripper
	token string
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.token)
	if t.base == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.base.RoundTrip(req)
}

// GetAvailableModels 获取可用的模型列表
func (c *Client) GetAvailableModels() []string {
	return []string{
		"qwen3-vl-plus",               // 图像分析主模型
		"qwen-flash",                  // 文本生成主模型
		"qwen3-max",                   // 复杂推理主模型
		"qwen3-omni-flash",            // 语音交互主模型
		"qwen3-vl-235b-a22b-instruct", // 图像分析备用模型
		"qwen-turbo",                  // 文本生成备用模型
		"qwen-max",                    // 复杂推理备用模型
	}
}

// ValidateModel 检查模型是否可用
func (c *Client) ValidateModel(model string) bool {
	availableModels := c.GetAvailableModels()
	for _, availableModel := range availableModels {
		if availableModel == model {
			return true
		}
	}
	return false
}

// GetModelForTask 根据任务类型推荐模型
func GetModelForTask(task string) string {
	switch task {
	case "image_analysis":
		return ModelImageAnalysis // qwen3-vl-plus - 视觉理解
	case "text_generation":
		return ModelTextGeneration // qwen-flash - 快速文本生成
	case "advanced_reasoning":
		return ModelAdvancedReasoning // qwen3-max - 复杂推理
	case "voice_interaction":
		return ModelVoiceInteraction // qwen3-omni-flash - 语音交互
	default:
		return ModelTextGeneration // 默认使用通用模型
	}
}

// GetModelCapabilities 获取模型能力说明
func GetModelCapabilities() map[string]string {
	return map[string]string{
		"qwen3-vl-plus":    "视觉理解，支持思考模式，图像分析最优，支持超长视频理解",
		"qwen-flash":       "思考+非思考模式融合，复杂推理优秀，指令遵循强",
		"qwen3-max":        "智能体编程优化，工具调用增强，领域SOTA水平",
		"qwen3-omni-flash": "多模态实时交互，支持文本、图像、音频、视频，119种语言文本交互，20种语言语音交互",
	}
}

// AnalyzeImage 分析图片
func (c *Client) AnalyzeImage(ctx context.Context, imageURL, prompt string) (*ImageAnalysisResult, error) {
	// 构建包含图片URL的prompt
	// 注意：如果内部AI服务支持多模态，可能需要使用不同的API格式
	// 这里先使用文本格式，将图片URL包含在prompt中
	fullPrompt := fmt.Sprintf("%s\n\n图片URL: %s", prompt, imageURL)

	req := openai.ChatCompletionRequest{
		Model: ModelImageAnalysis,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fullPrompt,
			},
		},
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Qwen API调用失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("Qwen API返回结果为空")
	}

	result := &ImageAnalysisResult{
		ObjectName: resp.Choices[0].Message.Content,
		// 这里可以根据实际需求解析更多信息
	}

	return result, nil
}

// GenerateQuestions 生成引导问题
func (c *Client) GenerateQuestions(ctx context.Context, contextInfo string, category string) ([]Question, error) {

	prompt := fmt.Sprintf(`基于以下信息为孩子生成3个引导性的探索问题：

上下文信息：%s
探索类别：%s

要求：
1. 问题要适合儿童理解
2. 问题要激发好奇心和思考
3. 问题难度要循序渐进（从简单到深入）
4. 每个问题都要有明确的类型标注
5. 确保所有内容适合儿童教育场景，避免任何不适宜内容

请以JSON格式返回，包含以下字段：
- content: 问题内容
- type: 问题类型（observation观察, reasoning推理, experiment实验, comparison比较）
- difficulty: 难度（basic基本, intermediate中级, advanced高级）
- purpose: 问题目的说明`, contextInfo, category)

	req := openai.ChatCompletionRequest{
		Model: ModelTextGeneration,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("生成问题失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("Qwen API返回结果为空")
	}

	// 这里需要解析JSON响应并转换为Question结构体
	// 为了简化，先返回空结果
	return []Question{}, nil
}

// PolishNote AI润色笔记
func (c *Client) PolishNote(ctx context.Context, rawContent, contextInfo string) (*PolishedNote, error) {

	prompt := fmt.Sprintf(`请帮孩子润色他的探索笔记，让它更清晰、有逻辑性。

原始内容：%s

上下文信息：%s

要求：
1. 保持孩子的原意和语言特色
2. 让表达更清晰准确
3. 添加适当的科学概念解释
4. 指出可能的疑问和下一步探索方向
5. 确保所有内容适合儿童教育场景，避免任何不适宜内容

请以JSON格式返回包含以下字段的结果：
- title: 笔记标题
- summary: 内容总结
- key_points: 关键要点数组
- scientific_concepts: 科学概念数组
- questions: 引发的疑问数组
- connections: 相关知识连接数组
- formatted_text: 格式化的文本内容`, rawContent, contextInfo)

	req := openai.ChatCompletionRequest{
		Model: ModelTextGeneration,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("润色笔记失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("Qwen API返回结果为空")
	}

	// 这里需要解析JSON响应并转换为PolishedNote结构体
	// 为了简化，先返回基本结果
	result := &PolishedNote{
		FormattedText: resp.Choices[0].Message.Content,
	}

	return result, nil
}

// GenerateReport 生成研究报告
func (c *Client) GenerateReport(ctx context.Context, projectData string) (*ResearchReport, error) {

	prompt := fmt.Sprintf(`基于孩子的研究数据生成一份研究报告：

项目数据：%s

请生成包含以下部分的研究报告：
1. 标题
2. 摘要
3. 引言
4. 方法论
5. 发现与结果
6. 讨论
7. 结论
8. 参考资料
9. 孩子独特见解
10. 下一步探索建议

重要要求：
- 确保所有内容适合儿童教育场景
- 避免任何不适宜的敏感内容
- 保持积极正面的教育导向

请以JSON格式返回报告内容。`, projectData)

	req := openai.ChatCompletionRequest{
		Model: ModelAdvancedReasoning,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("生成报告失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("Qwen API返回结果为空")
	}

	// 这里需要解析JSON响应并转换为ResearchReport结构体
	// 为了简化，先返回基本结果
	result := &ResearchReport{
		Content: resp.Choices[0].Message.Content,
	}

	return result, nil
}

// 数据结构定义

type ImageAnalysisResult struct {
	ObjectName     string   `json:"object_name"`
	Category       string   `json:"category"`
	Confidence     float64  `json:"confidence"`
	Description    string   `json:"description"`
	KeyFeatures    []string `json:"key_features"`
	ScientificName string   `json:"scientific_name"`
}

type Question struct {
	Content    string `json:"content"`
	Type       string `json:"type"`
	Difficulty string `json:"difficulty"`
	Purpose    string `json:"purpose"`
}

type PolishedNote struct {
	Title              string   `json:"title"`
	Summary            string   `json:"summary"`
	KeyPoints          []string `json:"key_points"`
	ScientificConcepts []string `json:"scientific_concepts"`
	Questions          []string `json:"questions"`
	Connections        []string `json:"connections"`
	FormattedText      string   `json:"formatted_text"`
}

type ResearchReport struct {
	Title         string      `json:"title"`
	Abstract      string      `json:"abstract"`
	Introduction  string      `json:"introduction"`
	Methodology   string      `json:"methodology"`
	Findings      []Finding   `json:"findings"`
	Discussion    string      `json:"discussion"`
	Conclusion    string      `json:"conclusion"`
	References    []Reference `json:"references"`
	ChildInsights string      `json:"child_insights"`
	NextSteps     []string    `json:"next_steps"`
	Content       string      `json:"content"` // 简化字段，用于存储完整内容
}

type Finding struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Evidence     []string `json:"evidence"`
	Significance string   `json:"significance"`
}

type Reference struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	URL    string `json:"url,omitempty"`
	Credit string `json:"credit"`
}
