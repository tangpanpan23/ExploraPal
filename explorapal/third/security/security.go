package security

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// SecurityClient 集团安全中心过滤服务客户端
type SecurityClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// Config 安全服务配置
type Config struct {
	BaseURL    string `json:"baseURL"`    // 安全中心服务地址
	APIKey     string `json:"apiKey"`     // 安全中心API密钥
	Timeout    int    `json:"timeout"`    // 超时时间(秒)
}

// NewSecurityClient 创建安全过滤客户端
func NewSecurityClient(config *Config) *SecurityClient {
	return &SecurityClient{
		baseURL: strings.TrimSuffix(config.BaseURL, "/"),
		apiKey:  config.APIKey,
		httpClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}
}

// ContentCheckRequest 内容检查请求
type ContentCheckRequest struct {
	Content     string   `json:"content"`      // 待检查内容
	ContentType string   `json:"content_type"` // 内容类型：text, image_url, audio_url
	Scene       string   `json:"scene"`        // 业务场景：education, children
	UserID      string   `json:"user_id"`      // 用户ID
	SessionID   string   `json:"session_id"`   // 会话ID
	CheckTypes  []string `json:"check_types"`  // 检查类型：politics, violence, minors, sensitive
}

// ContentCheckResponse 内容检查响应
type ContentCheckResponse struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	Data      ContentCheckResult     `json:"data"`
	RequestID string                 `json:"request_id"`
}

// ContentCheckResult 检查结果
type ContentCheckResult struct {
	Passed     bool                   `json:"passed"`      // 是否通过检查
	RiskLevel  string                 `json:"risk_level"`  // 风险等级：low, medium, high, reject
	Suggestion string                 `json:"suggestion"`  // 处理建议
	Details    []RiskDetail           `json:"details"`     // 风险详情
	FilteredContent string            `json:"filtered_content,omitempty"` // 过滤后的内容
}

// RiskDetail 风险详情
type RiskDetail struct {
	Type        string  `json:"type"`         // 风险类型
	Description string  `json:"description"`  // 风险描述
	Confidence  float64 `json:"confidence"`   // 置信度
	Suggestion  string  `json:"suggestion"`   // 处理建议
}

// CheckContent 检查内容合规性
func (c *SecurityClient) CheckContent(ctx context.Context, req *ContentCheckRequest) (*ContentCheckResult, error) {
	// 设置默认检查类型（面向儿童教育场景的重点风险）
	if len(req.CheckTypes) == 0 {
		req.CheckTypes = []string{"politics", "violence", "minors", "sensitive", "pornography"}
	}

	// 设置默认场景
	if req.Scene == "" {
		req.Scene = "children_education"
	}

	// 构建请求
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}

	// 发送HTTP请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST",
		c.baseURL+"/api/v1/content/check", strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("X-Client", "explorapal-security-client")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("请求安全中心失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("安全中心响应异常: %d", resp.StatusCode)
	}

	// 解析响应
	var checkResp ContentCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&checkResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	if checkResp.Code != 200 {
		return nil, fmt.Errorf("安全检查失败: %s", checkResp.Message)
	}

	return &checkResp.Data, nil
}

// FilterContent 过滤内容（如果内容不合规，返回过滤后的版本）
func (c *SecurityClient) FilterContent(ctx context.Context, content, contentType, userID, sessionID string) (string, error) {
	req := &ContentCheckRequest{
		Content:     content,
		ContentType: contentType,
		UserID:      userID,
		SessionID:   sessionID,
	}

	result, err := c.CheckContent(ctx, req)
	if err != nil {
		return "", err
	}

	if !result.Passed {
		// 如果有过滤后的内容，返回过滤版本；否则返回空字符串表示拒绝
		if result.FilteredContent != "" {
			return result.FilteredContent, nil
		}
		return "", fmt.Errorf("内容不合规，风险等级: %s, 建议: %s", result.RiskLevel, result.Suggestion)
	}

	return content, nil
}

// BatchCheckContent 批量检查内容
func (c *SecurityClient) BatchCheckContent(ctx context.Context, contents []string, contentType, userID, sessionID string) ([]ContentCheckResult, error) {
	var results []ContentCheckResult

	for _, content := range contents {
		req := &ContentCheckRequest{
			Content:     content,
			ContentType: contentType,
			UserID:      userID,
			SessionID:   sessionID,
		}

		result, err := c.CheckContent(ctx, req)
		if err != nil {
			// 单个检查失败时，记录错误但继续检查其他内容
			results = append(results, ContentCheckResult{
				Passed:    false,
				RiskLevel: "error",
				Suggestion: fmt.Sprintf("检查失败: %v", err),
			})
			continue
		}

		results = append(results, *result)
	}

	return results, nil
}

// IsContentSafe 快速检查内容是否安全（简化版）
func (c *SecurityClient) IsContentSafe(ctx context.Context, content, contentType, userID string) (bool, string, error) {
	req := &ContentCheckRequest{
		Content:     content,
		ContentType: contentType,
		UserID:      userID,
		CheckTypes:  []string{"politics", "violence", "minors", "sensitive"},
	}

	result, err := c.CheckContent(ctx, req)
	if err != nil {
		return false, "", err
	}

	suggestion := ""
	if len(result.Details) > 0 {
		suggestion = result.Details[0].Suggestion
	}

	return result.Passed, suggestion, nil
}
