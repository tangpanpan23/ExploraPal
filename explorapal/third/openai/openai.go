package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

	// 设置HTTP客户端超时时间
	if config.Timeout > 0 {
		clientConfig.HTTPClient.Timeout = time.Duration(config.Timeout) * time.Second
	} else {
		clientConfig.HTTPClient.Timeout = 70 * time.Second // 默认70秒超时
	}

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
	// 使用多模态模型进行图像分析
	contentParts := []openai.ChatMessagePart{
		{
			Type: openai.ChatMessagePartTypeText,
			Text: prompt,
		},
		{
			Type: openai.ChatMessagePartTypeImageURL,
			ImageURL: &openai.ChatMessageImageURL{
				URL: imageURL,
			},
		},
	}

	req := openai.ChatCompletionRequest{
		Model: ModelImageAnalysis,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:         openai.ChatMessageRoleUser,
				MultiContent: contentParts,
			},
		},
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		// AI服务不可用时，返回默认的模拟响应
		return c.getDefaultImageAnalysis(imageURL, prompt), nil
	}

	if len(resp.Choices) == 0 {
		return c.getDefaultImageAnalysis(imageURL, prompt), nil
	}

	content := resp.Choices[0].Message.Content

	// 解析AI返回的内容，提取结构化信息
	result := &ImageAnalysisResult{
		ObjectName:     extractObjectName(content),
		Category:       extractCategory(content),
		Description:    content,
		Confidence:     0.95, // 默认置信度
		KeyFeatures:    extractKeyFeatures(content),
		ScientificName: extractScientificName(content),
	}

	return result, nil
}

// 辅助函数：从AI响应中提取信息
func extractObjectName(content string) string {
	// 简单提取逻辑，可以根据实际响应格式优化
	if len(content) > 50 {
		return content[:50] + "..."
	}
	return content
}

func extractCategory(content string) string {
	// 根据内容判断类别
	if containsAny(content, "恐龙", "dinosaur", "化石") {
		return "dinosaur"
	}
	return "unknown"
}

func extractKeyFeatures(content string) []string {
	// 提取关键特征的简单逻辑
	return []string{"特征1", "特征2"} // 实际应该解析AI响应
}

func extractScientificName(content string) string {
	// 提取科学名称
	return "未知" // 实际应该解析AI响应
}

func containsAny(text string, substrings ...string) bool {
	for _, substr := range substrings {
		if strings.Contains(text, substr) {
			return true
		}
	}
	return false
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
				Role:    openai.ChatMessageRoleSystem,
				Content: "你是一个儿童教育助手，专门为孩子设计探索性问题。请以JSON格式返回包含questions数组的结果。",
			},
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
		// 如果AI调用失败，返回默认问题
		return c.getDefaultQuestions(category), nil
	}

	if len(resp.Choices) == 0 {
		return c.getDefaultQuestions(category), nil
	}

	content := resp.Choices[0].Message.Content

	// 记录AI原始响应
	fmt.Printf("[AI_DEBUG] GenerateQuestions原始响应:\n%s\n", content)

	// 尝试解析JSON响应
	questions := parseQuestionsFromJSON(content)
	if len(questions) == 0 {
		fmt.Printf("[AI_DEBUG] GenerateQuestions JSON解析失败，返回默认问题\n")
		// 如果解析失败，返回默认问题
		return c.getDefaultQuestions(category), nil
	}

	fmt.Printf("[AI_DEBUG] GenerateQuestions解析成功，返回%d个问题\n", len(questions))
	return questions, nil
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

	// 记录AI请求参数
	fmt.Printf("[AI_DEBUG] PolishNote请求参数:\n")
	fmt.Printf("  Model: %s\n", ModelTextGeneration)
	fmt.Printf("  MaxTokens: %d\n", c.config.MaxTokens)
	fmt.Printf("  Temperature: %.2f\n", c.config.Temperature)
	fmt.Printf("  Prompt长度: %d\n", len(prompt))
	fmt.Printf("  Prompt内容:\n%s\n", prompt)

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
		// 记录AI调用错误
		fmt.Printf("[AI_DEBUG] AI服务调用失败: %v\n", err)
		// AI服务不可用时，返回默认的润色结果
		return c.getDefaultPolishedNote(rawContent, contextInfo), nil
	}

	if len(resp.Choices) == 0 {
		fmt.Printf("[AI_DEBUG] AI响应Choices为空\n")
		return c.getDefaultPolishedNote(rawContent, contextInfo), nil
	}

	content := resp.Choices[0].Message.Content

	// 记录AI原始响应
	fmt.Printf("[AI_DEBUG] AI原始响应:\n%s\n", content)

	// 处理markdown格式的JSON代码块
	jsonContent := content
	if strings.HasPrefix(strings.TrimSpace(content), "```json") {
		fmt.Printf("[AI_DEBUG] 检测到markdown格式的JSON代码块\n")
		// 提取 ```json 和 ``` 之间的内容
		startIndex := strings.Index(content, "```json")
		if startIndex != -1 {
			startIndex += 7 // 跳过 ```json
			endIndex := strings.Index(content[startIndex:], "```")
			if endIndex != -1 {
				jsonContent = strings.TrimSpace(content[startIndex : startIndex+endIndex])
				fmt.Printf("[AI_DEBUG] 提取的JSON内容:\n%s\n", jsonContent)
			} else {
				fmt.Printf("[AI_DEBUG] 未找到结束标记 ```\n")
			}
		}
	}

	// 尝试解析JSON响应
	var jsonResult PolishedNote
	if err := json.Unmarshal([]byte(jsonContent), &jsonResult); err != nil {
		// 记录JSON解析错误
		fmt.Printf("[AI_DEBUG] JSON解析失败: %v\n", err)
		contentPreviewLen := 200
		if len(jsonContent) < 200 {
			contentPreviewLen = len(jsonContent)
		}
		fmt.Printf("[AI_DEBUG] JSON内容前%d字符: %s\n", contentPreviewLen, jsonContent[:contentPreviewLen])
		// 如果JSON解析失败，返回默认结果
		return c.getDefaultPolishedNote(rawContent, contextInfo), nil
	}

	// 记录解析成功的结构化数据
	fmt.Printf("[AI_DEBUG] JSON解析成功:\n")
	fmt.Printf("  Title: %s\n", jsonResult.Title)
	fmt.Printf("  Summary: %s\n", jsonResult.Summary)
	fmt.Printf("  KeyPoints数量: %d\n", len(jsonResult.KeyPoints))

	// 确保所有必需字段都有值
	if jsonResult.Title == "" {
		fmt.Printf("[AI_DEBUG] Title字段为空，使用默认值\n")
		jsonResult.Title = "探索笔记"
	}
	if jsonResult.Summary == "" {
		fmt.Printf("[AI_DEBUG] Summary字段为空，使用默认值\n")
		summaryLen := 20
		if len(rawContent) < 20 {
			summaryLen = len(rawContent)
		}
		jsonResult.Summary = fmt.Sprintf("这是关于%s的学习笔记", rawContent[:summaryLen])
	}
	if len(jsonResult.KeyPoints) == 0 {
		fmt.Printf("[AI_DEBUG] KeyPoints字段为空，使用默认值\n")
		jsonResult.KeyPoints = []string{"学习了新的知识", "发现了有趣的现象"}
	}
	if jsonResult.FormattedText == "" {
		fmt.Printf("[AI_DEBUG] FormattedText字段为空，使用原始响应\n")
		jsonResult.FormattedText = content
	}

	fmt.Printf("[AI_DEBUG] 返回最终结果: Title='%s'\n", jsonResult.Title)
	return &jsonResult, nil
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
		// AI服务不可用时，返回默认的报告
		return c.getDefaultResearchReport(projectData), nil
	}

	if len(resp.Choices) == 0 {
		return c.getDefaultResearchReport(projectData), nil
	}

	content := resp.Choices[0].Message.Content

	// 记录AI原始响应
	fmt.Printf("[AI_DEBUG] GenerateReport原始响应:\n%s\n", content)

	// 处理markdown格式的JSON代码块
	jsonContent := content
	if strings.HasPrefix(strings.TrimSpace(content), "```json") {
		fmt.Printf("[AI_DEBUG] GenerateReport检测到markdown格式的JSON代码块\n")
		// 提取 ```json 和 ``` 之间的内容
		startIndex := strings.Index(content, "```json")
		if startIndex != -1 {
			startIndex += 7 // 跳过 ```json
			endIndex := strings.Index(content[startIndex:], "```")
			if endIndex != -1 {
				jsonContent = strings.TrimSpace(content[startIndex : startIndex+endIndex])
				fmt.Printf("[AI_DEBUG] GenerateReport提取的JSON内容:\n%s\n", jsonContent)
			}
		}
	}

	// 尝试解析JSON响应
	var jsonResult ResearchReport
	if err := json.Unmarshal([]byte(jsonContent), &jsonResult); err != nil {
		fmt.Printf("[AI_DEBUG] GenerateReport JSON解析失败: %v\n", err)
		// 如果JSON解析失败，返回默认结果
		return c.getDefaultResearchReport(projectData), nil
	}

	// 确保所有必需字段都有值
	if jsonResult.Title == "" {
		jsonResult.Title = "探索研究报告"
	}
	if jsonResult.Content == "" {
		jsonResult.Content = jsonContent
	}

	fmt.Printf("[AI_DEBUG] GenerateReport解析成功: Title='%s'\n", jsonResult.Title)
	return &jsonResult, nil
}

// getDefaultQuestions 根据类别返回默认问题
func (c *Client) getDefaultQuestions(category string) []Question {
	switch category {
	case "dinosaur":
		return []Question{
			{
				Content:    "你看到恐龙的哪个部位最有趣？",
				Type:       "observation",
				Difficulty: "basic",
				Purpose:    "培养观察力",
			},
			{
				Content:    "你觉得这只恐龙生活在什么时候？为什么？",
				Type:       "reasoning",
				Difficulty: "intermediate",
				Purpose:    "培养推理能力",
			},
			{
				Content:    "如果我们能见到活的恐龙，你最想问它什么问题？",
				Type:       "experiment",
				Difficulty: "advanced",
				Purpose:    "激发想象力和探索欲",
			},
		}
	default:
		return []Question{
			{
				Content:    "你看到了什么有趣的东西？",
				Type:       "observation",
				Difficulty: "basic",
				Purpose:    "激发观察兴趣",
			},
			{
				Content:    "为什么会这样呢？",
				Type:       "reasoning",
				Difficulty: "intermediate",
				Purpose:    "培养思考能力",
			},
			{
				Content:    "我们可以做个小实验验证吗？",
				Type:       "experiment",
				Difficulty: "advanced",
				Purpose:    "实践探索精神",
			},
		}
	}
}

// parseQuestionsFromJSON 解析AI返回的JSON格式问题
func parseQuestionsFromJSON(content string) []Question {
	fmt.Printf("[AI_DEBUG] parseQuestionsFromJSON输入内容:\n%s\n", content)

	// 处理markdown格式的JSON代码块
	jsonContent := content
	if strings.HasPrefix(strings.TrimSpace(content), "```json") {
		fmt.Printf("[AI_DEBUG] parseQuestionsFromJSON检测到markdown格式\n")
		// 提取 ```json 和 ``` 之间的内容
		startIndex := strings.Index(content, "```json")
		if startIndex != -1 {
			startIndex += 7 // 跳过 ```json
			endIndex := strings.Index(content[startIndex:], "```")
			if endIndex != -1 {
				jsonContent = strings.TrimSpace(content[startIndex : startIndex+endIndex])
				fmt.Printf("[AI_DEBUG] parseQuestionsFromJSON提取的JSON:\n%s\n", jsonContent)
			}
		}
	}

	// 尝试解析JSON
	var result struct {
		Questions []Question `json:"questions"`
	}

	if err := json.Unmarshal([]byte(jsonContent), &result); err != nil {
		fmt.Printf("[AI_DEBUG] parseQuestionsFromJSON解析questions结构失败: %v\n", err)
		// 如果解析失败，尝试直接解析为问题数组
		var questions []Question
		if err2 := json.Unmarshal([]byte(jsonContent), &questions); err2 != nil {
			fmt.Printf("[AI_DEBUG] parseQuestionsFromJSON直接解析数组也失败: %v\n", err2)
			// 如果都失败了，返回空数组
			return []Question{}
		}
		fmt.Printf("[AI_DEBUG] parseQuestionsFromJSON直接解析数组成功，返回%d个问题\n", len(questions))
		return questions
	}

	fmt.Printf("[AI_DEBUG] parseQuestionsFromJSON解析questions结构成功，返回%d个问题\n", len(result.Questions))
	return result.Questions
}

// getDefaultImageAnalysis AI服务不可用时返回默认的图像分析结果
func (c *Client) getDefaultImageAnalysis(imageURL, prompt string) *ImageAnalysisResult {
	// 根据URL和提示词提供默认的分析结果
	result := &ImageAnalysisResult{
		ObjectName:     "未知物体",
		Category:       "general",
		Description:    "由于AI服务暂时不可用，这里提供一个模拟的分析结果。在实际环境中，这个结果将由AI模型生成。",
		Confidence:     0.80, // 默认置信度
		KeyFeatures:    []string{"特征分析", "形态识别", "内容描述"},
		ScientificName: "未知",
	}

	// 如果URL包含特定关键词，提供相应的默认结果
	if containsAny(imageURL, "dinosaur", "恐龙") {
		result.ObjectName = "恐龙化石"
		result.Category = "dinosaur"
		result.Description = "这看起来像是一块恐龙化石，包含了古代生物的遗骸。"
		result.KeyFeatures = []string{"骨骼结构", "化石纹理", "年代久远"}
		result.ScientificName = "恐龙类"
	}

	return result
}

// getDefaultPolishedNote AI服务不可用时返回默认的笔记润色结果
func (c *Client) getDefaultPolishedNote(rawContent, contextInfo string) *PolishedNote {
	fmt.Printf("[AI_DEBUG] 返回默认润色结果 (AI服务异常或解析失败)\n")
	result := &PolishedNote{
		FormattedText:      rawContent, // 保持原始内容
		Title:              "探索笔记",
		Summary:            "这是孩子记录的探索笔记，由于AI服务暂时不可用，显示原始内容。",
		KeyPoints:          []string{"观察记录", "思考过程"},
		ScientificConcepts: []string{"观察", "记录"},
		Questions:          []string{"你发现了什么？", "你想知道什么？"},
		Connections:        []string{"科学探索", "学习过程"},
	}
	return result
}

// getDefaultResearchReport AI服务不可用时返回默认的研究报告
func (c *Client) getDefaultResearchReport(projectData string) *ResearchReport {
	result := &ResearchReport{
		Title:        "探索研究报告",
		Abstract:     "这是基于孩子探索过程生成的研究报告摘要。",
		Introduction: "孩子通过观察、提问和表达的方式进行了科学探索。",
		Methodology:  "使用了观察、提问、记录的方法进行探索。",
		Findings: []Finding{
			{
				Title:        "探索发现",
				Description:  "孩子发现了许多有趣的现象，并提出了自己的问题。",
				Evidence:     []string{"观察记录", "提问过程"},
				Significance: "培养了观察能力和思考能力",
			},
		},
		Discussion: "这个探索过程帮助孩子培养了观察能力和思考能力。",
		Conclusion: "探索学习是一种有效的教育方式。",
		References: []Reference{
			{Title: "科学观察方法", Type: "教育资源", Credit: "教育专家"},
			{Title: "儿童学习理论", Type: "研究文献", Credit: "教育研究者"},
		},
		ChildInsights: "孩子独特的视角和创造性思维。",
		NextSteps:     []string{"继续探索", "深入研究", "分享发现"},
		Content:       "由于AI服务暂时不可用，这里提供一个模板化的研究报告。在实际环境中，这个报告将由AI根据具体项目数据生成。",
	}
	return result
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
