package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// Client OpenAI客户端
type Client struct {
	client *openai.Client
	config *Config
}

// Config 配置
type Config struct {
	APIKey     string  `json:"apiKey"`
	BaseURL    string  `json:"baseURL,omitempty"`
	Timeout    int     `json:"timeout,omitempty"`    // 超时时间(秒)
	MaxTokens  int     `json:"maxTokens,omitempty"`  // 最大token数
	Temperature float32 `json:"temperature,omitempty"` // 温度参数
}

// NewClient 创建OpenAI客户端
func NewClient(config *Config) *Client {
	clientConfig := openai.DefaultConfig(config.APIKey)
	if config.BaseURL != "" {
		clientConfig.BaseURL = config.BaseURL
	}

	return &Client{
		client: openai.NewClientWithConfig(clientConfig),
		config: config,
	}
}

// AnalyzeImage 分析图片
func (c *Client) AnalyzeImage(ctx context.Context, imageURL, prompt string) (*ImageAnalysisResult, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4VisionPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: []openai.ChatMessagePart{
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
				},
			},
		},
		MaxTokens: c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("OpenAI API调用失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("OpenAI API返回结果为空")
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

请以JSON格式返回，包含以下字段：
- content: 问题内容
- type: 问题类型（observation观察, reasoning推理, experiment实验, comparison比较）
- difficulty: 难度（basic基本, intermediate中级, advanced高级）
- purpose: 问题目的说明`, contextInfo, category)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens: c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("生成问题失败: %w", err)
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

请以JSON格式返回包含以下字段的结果：
- title: 笔记标题
- summary: 内容总结
- key_points: 关键要点数组
- scientific_concepts: 科学概念数组
- questions: 引发的疑问数组
- connections: 相关知识连接数组
- formatted_text: 格式化的文本内容`, rawContent, contextInfo)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens: c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("润色笔记失败: %w", err)
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

请以JSON格式返回报告内容。`, projectData)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens: c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("生成报告失败: %w", err)
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
	Title             string   `json:"title"`
	Summary           string   `json:"summary"`
	KeyPoints         []string `json:"key_points"`
	ScientificConcepts []string `json:"scientific_concepts"`
	Questions         []string `json:"questions"`
	Connections       []string `json:"connections"`
	FormattedText     string   `json:"formatted_text"`
}

type ResearchReport struct {
	Title         string     `json:"title"`
	Abstract      string     `json:"abstract"`
	Introduction  string     `json:"introduction"`
	Methodology   string     `json:"methodology"`
	Findings      []Finding  `json:"findings"`
	Discussion    string     `json:"discussion"`
	Conclusion    string     `json:"conclusion"`
	References    []Reference `json:"references"`
	ChildInsights string     `json:"child_insights"`
	NextSteps     []string   `json:"next_steps"`
	Content       string     `json:"content"` // 简化字段，用于存储完整内容
}

type Finding struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Evidence    []string `json:"evidence"`
	Significance string  `json:"significance"`
}

type Reference struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	URL     string `json:"url,omitempty"`
	Credit  string `json:"credit"`
}
