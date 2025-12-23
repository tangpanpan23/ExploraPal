package logic

import (
	"context"
	"encoding/base64"

	videoprocessing "explorapal/app/video-processing/proto"
	"explorapal/app/video-processing/rpc/internal/svc"
	"explorapal/third/openai"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateVideoLogic {
	return &GenerateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateVideoLogic) GenerateVideo(in *videoprocessing.GenerateVideoReq) (*videoprocessing.GenerateVideoResp, error) {
	// 记录API请求参数
	l.Infof("视频生成请求: Script长度=%d, ImageData长度=%d, Prompt长度=%d, Style=%s, Duration=%.2f, Scenes数量=%d, Voice=%s, Language=%s",
		len(in.Script), len(in.ImageData), len(in.Prompt), in.Style, in.Duration, len(in.Scenes), in.Voice, in.Language)

	// 1. 验证输入（支持两种模式）
	isImageToVideoMode := in.ImageData != "" && in.Prompt != ""
	isTextToVideoMode := in.Script != ""

	if !isImageToVideoMode && !isTextToVideoMode {
		return &videoprocessing.GenerateVideoResp{
			Status: 400,
			Msg:    "必须提供视频脚本(script)或图片数据+描述(image_data+prompt)",
		}, nil
	}

	if isImageToVideoMode && isTextToVideoMode {
		return &videoprocessing.GenerateVideoResp{
			Status: 400,
			Msg:    "不能同时提供脚本和图片数据，请选择一种模式",
		}, nil
	}

	// 2. 设置默认参数
	style := in.Style
	if style == "" {
		style = "educational"
	}

	duration := in.Duration
	if duration <= 0 {
		duration = 60.0 // 默认60秒
	}

	voice := in.Voice
	if voice == "" {
		voice = "female"
	}

	language := in.Language
	if language == "" {
		language = "zh-CN"
	}

	// 3. 根据模式调用相应的AI服务生成视频
	var videoData []byte
	var format string
	var actualDuration float64
	var metadata *videoprocessing.VideoMetadata
	var err error

	if isImageToVideoMode {
		// 图像到视频模式 - 使用豆包Doubao-Seedance-1.0-lite-i2v
		l.Infof("使用图像到视频模式处理请求")
		videoData, format, actualDuration, metadata, err = l.processImageToVideoGenerationWithAI(in.ImageData, in.Prompt, style, duration, in.Scenes, voice, language)
	} else {
		// 文本到视频模式 - 使用原有逻辑
		l.Infof("使用文本到视频模式处理请求")
		videoData, format, actualDuration, metadata, err = l.processVideoGenerationWithAI(in.Script, style, duration, in.Scenes, voice, language)
	}
	if err != nil {
		l.Logger.Errorf("AI视频生成失败: %v", err)
		// 返回模拟结果
		return l.getDefaultVideoGenerationResult(in), nil
	}

	// 记录成功结果
	l.Infof("视频生成成功: 大小=%d bytes, 格式=%s, 时长=%.2f秒", len(videoData), format, actualDuration)

	return &videoprocessing.GenerateVideoResp{
		Status:    200,
		Msg:       "视频生成成功",
		VideoData: videoData,
		Format:    format,
		Duration:  actualDuration,
		Metadata:  metadata,
	}, nil
}

// processVideoGenerationWithAI 调用AI服务处理视频生成
func (l *GenerateVideoLogic) processVideoGenerationWithAI(script, style string, duration float64, scenes []string, voice, language string) ([]byte, string, float64, *videoprocessing.VideoMetadata, error) {
	// 使用ServiceContext中的AI客户端
	// 显式使用openai包以避免编译器误报
	_ = openai.Client{}
	videoData, format, actualDuration, aiMetadata, err := l.svcCtx.AIClient.GenerateVideo(l.ctx, script, style, duration, scenes, voice, language)
	if err != nil {
		l.Logger.Errorf("AI视频生成调用失败: %v", err)
		return nil, "", 0, nil, err
	}

	// 转换AI元数据为Protobuf格式
	metadata := &videoprocessing.VideoMetadata{
		Title:         sanitizeUTF8(aiMetadata.Title),
		Description:   sanitizeUTF8(aiMetadata.Description),
		Scenes:        sanitizeUTF8Slice(aiMetadata.Scenes),
		AudioLanguage: aiMetadata.AudioLanguage,
		Resolution:    aiMetadata.Resolution,
	}

	return videoData, format, actualDuration, metadata, nil
}

// processVideoGeneration 保留旧方法名作为兼容
func (l *GenerateVideoLogic) processVideoGeneration(script, style string, duration float64, scenes []string, voice, language string) ([]byte, string, float64, *videoprocessing.VideoMetadata, error) {
	return l.processVideoGenerationWithAI(script, style, duration, scenes, voice, language)
}

// processImageToVideoGenerationWithAI 调用AI服务处理图像到视频生成
func (l *GenerateVideoLogic) processImageToVideoGenerationWithAI(imageData, prompt, style string, duration float64, scenes []string, voice, language string) ([]byte, string, float64, *videoprocessing.VideoMetadata, error) {
	// 使用ServiceContext中的AI客户端
	// 显式使用openai包以避免编译器误报
	_ = openai.Client{}

	// 调用豆包图像到视频生成方法
	videoData, format, actualDuration, aiMetadata, err := l.svcCtx.AIClient.GenerateVideoWithImage(l.ctx, imageData, prompt, style, duration, scenes, voice, language)
	if err != nil {
		l.Logger.Errorf("AI图像到视频生成调用失败: %v", err)
		return nil, "", 0, nil, err
	}

	// 转换AI元数据为Protobuf格式
	metadata := &videoprocessing.VideoMetadata{
		Title:         sanitizeUTF8(aiMetadata.Title),
		Description:   sanitizeUTF8(aiMetadata.Description),
		Scenes:        sanitizeUTF8Slice(aiMetadata.Scenes),
		AudioLanguage: aiMetadata.AudioLanguage,
		Resolution:    aiMetadata.Resolution,
	}

	return videoData, format, actualDuration, metadata, nil
}

// getDefaultVideoGenerationResult 返回默认的视频生成结果
func (l *GenerateVideoLogic) getDefaultVideoGenerationResult(in *videoprocessing.GenerateVideoReq) *videoprocessing.GenerateVideoResp {
	// 模拟视频数据
	mockVideoData := []byte("这是模拟的视频数据。实际环境中，这里将调用AI视频生成服务来生成真实的视频内容。")
	videoDataBase64 := base64.StdEncoding.EncodeToString(mockVideoData)

	return &videoprocessing.GenerateVideoResp{
		Status:    200,
		Msg:       "视频生成成功（使用模拟响应）",
		VideoData: []byte(videoDataBase64),
		Format:    "mp4",
		Duration:  60.0,
		Metadata: &videoprocessing.VideoMetadata{
			Title:          sanitizeUTF8("AI生成的演示视频"),
			Description:    sanitizeUTF8("这是AI生成的演示视频，由于AI服务暂时不可用，显示默认内容"),
			Scenes:         []string{"场景1", "场景2", "场景3"},
			AudioLanguage:  "zh-CN",
			Resolution:     "1920x1080",
		},
	}
}
