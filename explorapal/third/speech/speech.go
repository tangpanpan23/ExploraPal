package speech

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client 阿里云语音服务客户端
type Client struct {
	accessKeyId     string
	accessKeySecret string
	appKey          string
	host            string
	httpClient      *http.Client
}

// Config 语音服务配置
type Config struct {
	AccessKeyId     string `json:"accessKeyId"`     // 阿里云AccessKey ID
	AccessKeySecret string `json:"accessKeySecret"` // 阿里云AccessKey Secret
	AppKey          string `json:"appKey"`          // 语音服务AppKey
	Region          string `json:"region"`          // 地域，如"cn-shanghai"
}

// ASRRequest 语音识别请求
type ASRRequest struct {
	Format      string `json:"format"`       // 音频格式：wav, mp3, m4a等
	SampleRate  int    `json:"sample_rate"`  // 采样率：8000, 16000等
	Language    string `json:"language"`     // 语言：zh-CN, en-US等
	AudioBase64 string `json:"audio_base64"` // base64编码的音频数据
}

// ASRResponse 语音识别响应
type ASRResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      struct {
		Result string  `json:"result"`     // 识别结果
		Confidence float64 `json:"confidence"` // 置信度
	} `json:"data"`
}

// TTSRequest 语音合成请求
type TTSRequest struct {
	Text   string `json:"text"`             // 待合成的文本
	Voice  string `json:"voice,omitempty"`  // 音色，可选
	Format string `json:"format,omitempty"` // 输出格式：wav, mp3等
	SpeechRate int    `json:"speech_rate,omitempty"` // 语速：-500到500
	PitchRate  int    `json:"pitch_rate,omitempty"`  // 音调：-500到500
}

// TTSResponse 语音合成响应
type TTSResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      struct {
		AudioBase64 string `json:"audio_base64"` // base64编码的音频数据
		Format      string `json:"format"`       // 音频格式
		SampleRate  int    `json:"sample_rate"`  // 采样率
	} `json:"data"`
}

// NewClient 创建语音服务客户端
func NewClient(config *Config) *Client {
	host := fmt.Sprintf("https://nls-gateway-%s.aliyuncs.com", config.Region)

	return &Client{
		accessKeyId:     config.AccessKeyId,
		accessKeySecret: config.AccessKeySecret,
		appKey:          config.AppKey,
		host:            host,
		httpClient:      &http.Client{Timeout: 30 * time.Second},
	}
}

// SpeechToText 语音转文字
func (c *Client) SpeechToText(ctx context.Context, audioData []byte, format string, sampleRate int, language string) (string, error) {
	// 将音频数据编码为base64
	audioBase64 := base64.StdEncoding.EncodeToString(audioData)

	req := ASRRequest{
		Format:      format,
		SampleRate:  sampleRate,
		Language:    language,
		AudioBase64: audioBase64,
	}

	// 发送请求
	resp, err := c.doASRRequest(ctx, req)
	if err != nil {
		return "", fmt.Errorf("语音识别请求失败: %w", err)
	}

	if resp.Code != 200 {
		return "", fmt.Errorf("语音识别失败: %s", resp.Message)
	}

	return resp.Data.Result, nil
}

// TextToSpeech 文字转语音
func (c *Client) TextToSpeech(ctx context.Context, text string, voice string, format string) ([]byte, error) {
	req := TTSRequest{
		Text:   text,
		Voice:  voice,
		Format: format,
	}

	// 设置默认参数
	if req.Voice == "" {
		req.Voice = "xiaoyun" // 默认女声
	}
	if req.Format == "" {
		req.Format = "mp3"
	}

	// 发送请求
	resp, err := c.doTTSRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("语音合成请求失败: %w", err)
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("语音合成失败: %s", resp.Message)
	}

	// 解码base64音频数据
	audioData, err := base64.StdEncoding.DecodeString(resp.Data.AudioBase64)
	if err != nil {
		return nil, fmt.Errorf("音频数据解码失败: %w", err)
	}

	return audioData, nil
}

// doASRRequest 执行语音识别请求
func (c *Client) doASRRequest(ctx context.Context, req ASRRequest) (*ASRResponse, error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}

	// 构建请求URL
	requestURL := fmt.Sprintf("%s/stream/v1/asr", c.host)

	// 创建HTTP请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	c.setHeaders(httpReq, "asr")

	// 发送请求
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API响应异常: %d", resp.StatusCode)
	}

	// 解析响应
	var asrResp ASRResponse
	if err := json.NewDecoder(resp.Body).Decode(&asrResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &asrResp, nil
}

// doTTSRequest 执行语音合成请求
func (c *Client) doTTSRequest(ctx context.Context, req TTSRequest) (*TTSResponse, error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}

	// 构建请求URL
	requestURL := fmt.Sprintf("%s/stream/v1/tts", c.host)

	// 创建HTTP请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	c.setHeaders(httpReq, "tts")

	// 发送请求
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API响应异常: %d", resp.StatusCode)
	}

	// 解析响应
	var ttsResp TTSResponse
	if err := json.NewDecoder(resp.Body).Decode(&ttsResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &ttsResp, nil
}

// setHeaders 设置请求头
func (c *Client) setHeaders(req *http.Request, service string) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-NLS-App-Key", c.appKey)
	req.Header.Set("X-NLS-Timestamp", timestamp)
	req.Header.Set("X-NLS-Service-Type", service)

	// 生成签名
	signature := c.generateSignature(service, timestamp)
	req.Header.Set("X-NLS-Signature", signature)
}

// generateSignature 生成请求签名
func (c *Client) generateSignature(service, timestamp string) string {
	// 构建签名字符串
	signString := fmt.Sprintf("POST\napplication/json\n%s\n/stream/v1/%s", timestamp, service)

	// 使用HMAC-SHA256生成签名
	h := hmac.New(sha256.New, []byte(c.accessKeySecret))
	h.Write([]byte(signString))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}
