package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// MCPClient TALect MCP客户端
type MCPClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewMCPClient 创建MCP客户端
func NewMCPClient(baseURL string) *MCPClient {
	return &MCPClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// MCPRequest MCP请求结构
type MCPRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// MCPResponse MCP响应结构
type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
}

// MCPError MCP错误结构
type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// InitializeRequest 初始化请求
type InitializeRequest struct {
	ProtocolVersion string `json:"protocolVersion"`
	Capabilities    map[string]interface{} `json:"capabilities"`
	ClientInfo      map[string]interface{} `json:"clientInfo"`
}

// ToolsCallRequest 工具调用请求
type ToolsCallRequest struct {
	Name      string      `json:"name"`
	Arguments interface{} `json:"arguments"`
}

// ToolsCallResponse 工具调用响应
type ToolsCallResponse struct {
	Content []Content `json:"content"`
	IsError bool      `json:"isError,omitempty"`
}

// Content 内容结构
type Content struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

// SearchMaterialsParams 搜索素材参数
type SearchMaterialsParams struct {
	Query   string   `json:"query"`
	Grade   []string `json:"grade,omitempty"`
	Subject string   `json:"subject,omitempty"`
	Limit   int      `json:"limit,omitempty"`
}

// GetMaterialDetailParams 获取素材详情参数
type GetMaterialDetailParams struct {
	MaterialID string `json:"material_id"`
}

// GenerateLessonPlanParams 生成教案参数
type GenerateLessonPlanParams struct {
	MaterialIDs  []string `json:"material_ids"`
	Objectives   []string `json:"objectives"`
	Grade        string   `json:"grade"`
	StudentLevel string   `json:"student_level,omitempty"`
	Duration     int      `json:"duration,omitempty"`
}

// Initialize 初始化MCP连接
func (c *MCPClient) Initialize() error {
	req := MCPRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
		Params: InitializeRequest{
			ProtocolVersion: "2024-11-05",
			Capabilities: map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			ClientInfo: map[string]interface{}{
				"name":    "ExploraPal AI Learning Platform",
				"version": "1.0.0",
			},
		},
	}

	var resp MCPResponse
	if err := c.doRequest(req, &resp); err != nil {
		return fmt.Errorf("failed to initialize MCP client: %w", err)
	}

	if resp.Error != nil {
		return fmt.Errorf("MCP initialization error: %s", resp.Error.Message)
	}

	fmt.Println("✅ MCP客户端初始化成功")
	return nil
}

// SearchTeachingMaterials 搜索教学素材
func (c *MCPClient) SearchTeachingMaterials(query string, grade []string, subject string, limit int) (*ToolsCallResponse, error) {
	params := SearchMaterialsParams{
		Query:   query,
		Grade:   grade,
		Subject: subject,
		Limit:   limit,
	}

	req := MCPRequest{
		JSONRPC: "2.0",
		ID:      uuid.New().String(),
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name":      "search_teaching_materials",
			"arguments": params,
		},
	}

	var resp MCPResponse
	if err := c.doRequest(req, &resp); err != nil {
		return nil, fmt.Errorf("failed to search materials: %w", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("MCP search error: %s", resp.Error.Message)
	}

	// 解析响应
	var toolResp ToolsCallResponse
	if result, ok := resp.Result.(map[string]interface{}); ok {
		if content, ok := result["content"].([]interface{}); ok && len(content) > 0 {
			if contentMap, ok := content[0].(map[string]interface{}); ok {
				if text, ok := contentMap["text"].(string); ok {
					toolResp.Content = []Content{{Type: "text", Text: text}}
				}
			}
		}
	}

	return &toolResp, nil
}

// GetMaterialDetail 获取素材详情
func (c *MCPClient) GetMaterialDetail(materialID string) (*ToolsCallResponse, error) {
	params := GetMaterialDetailParams{
		MaterialID: materialID,
	}

	req := MCPRequest{
		JSONRPC: "2.0",
		ID:      uuid.New().String(),
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name":      "get_material_detail",
			"arguments": params,
		},
	}

	var resp MCPResponse
	if err := c.doRequest(req, &resp); err != nil {
		return nil, fmt.Errorf("failed to get material detail: %w", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("MCP get detail error: %s", resp.Error.Message)
	}

	// 解析响应
	var toolResp ToolsCallResponse
	if result, ok := resp.Result.(map[string]interface{}); ok {
		if content, ok := result["content"].([]interface{}); ok && len(content) > 0 {
			if contentMap, ok := content[0].(map[string]interface{}); ok {
				if text, ok := contentMap["text"].(string); ok {
					toolResp.Content = []Content{{Type: "text", Text: text}}
				}
			}
		}
	}

	return &toolResp, nil
}

// GenerateLessonPlan 生成教案
func (c *MCPClient) GenerateLessonPlan(materialIDs []string, objectives []string, grade string, studentLevel string, duration int) (*ToolsCallResponse, error) {
	params := GenerateLessonPlanParams{
		MaterialIDs:  materialIDs,
		Objectives:   objectives,
		Grade:        grade,
		StudentLevel: studentLevel,
		Duration:     duration,
	}

	req := MCPRequest{
		JSONRPC: "2.0",
		ID:      uuid.New().String(),
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name":      "generate_lesson_plan",
			"arguments": params,
		},
	}

	var resp MCPResponse
	if err := c.doRequest(req, &resp); err != nil {
		return nil, fmt.Errorf("failed to generate lesson plan: %w", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("MCP generate lesson plan error: %s", resp.Error.Message)
	}

	// 解析响应
	var toolResp ToolsCallResponse
	if result, ok := resp.Result.(map[string]interface{}); ok {
		if content, ok := result["content"].([]interface{}); ok && len(content) > 0 {
			if contentMap, ok := content[0].(map[string]interface{}); ok {
				if text, ok := contentMap["text"].(string); ok {
					toolResp.Content = []Content{{Type: "text", Text: text}}
				}
			}
		}
	}

	return &toolResp, nil
}

// doRequest 执行HTTP请求
func (c *MCPClient) doRequest(req MCPRequest, resp *MCPResponse) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error %d: %s", httpResp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}
