package logic

import (
	"context"

	videoprocessing "explorapal/app/video-processing/proto"
	"explorapal/app/video-processing/rpc/internal/svc"
	"explorapal/third/openai"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalyzeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalyzeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzeVideoLogic {
	return &AnalyzeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnalyzeVideoLogic) AnalyzeVideo(in *videoprocessing.AnalyzeVideoReq) (*videoprocessing.AnalyzeVideoResp, error) {
	// 记录API请求参数
	l.Infof("视频分析请求: VideoData大小=%d bytes, Format=%s, AnalysisType=%s, Duration=%.2f",
		len(in.VideoData), in.Format, in.AnalysisType, in.Duration)

	// 1. 验证输入
	if len(in.VideoData) == 0 {
		return &videoprocessing.AnalyzeVideoResp{
			Status: 400,
			Msg:    "视频数据不能为空",
		}, nil
	}

	// 2. 调用AI服务进行视频分析
	result, err := l.processVideoAnalysisWithAI(in)
	if err != nil {
		l.Logger.Errorf("AI视频分析失败: %v", err)
		// 返回模拟结果
		return l.getDefaultVideoAnalysisResult(in), nil
	}

	// 记录成功结果
	l.Infof("视频分析成功: 场景数量=%d, 物体数量=%d, 文字数量=%d", len(result.Scenes), len(result.Objects), len(result.Texts))

	return &videoprocessing.AnalyzeVideoResp{
		Status: 200,
		Msg:    "视频分析成功",
		Result: result,
	}, nil
}

// processVideoAnalysisWithAI 调用AI服务处理视频分析
func (l *AnalyzeVideoLogic) processVideoAnalysisWithAI(in *videoprocessing.AnalyzeVideoReq) (*videoprocessing.VideoAnalysisResult, error) {
	// 使用ServiceContext中的AI客户端
	// 显式使用openai包以避免编译器误报
	_ = openai.Client{}
	analysisType := in.AnalysisType
	if analysisType == "" {
		analysisType = "content"
	}

	aiResult, err := l.svcCtx.AIClient.AnalyzeVideo(l.ctx, in.VideoData, in.Format, analysisType, in.Duration)
	if err != nil {
		l.Logger.Errorf("AI视频分析调用失败: %v", err)
		return nil, err
	}

	// 转换AI结果为Protobuf格式
	result := &videoprocessing.VideoAnalysisResult{}

	// 转换场景分析
	if aiResult.Scenes != nil {
		result.Scenes = make([]*videoprocessing.SceneAnalysis, len(aiResult.Scenes))
		for i, scene := range aiResult.Scenes {
			result.Scenes[i] = &videoprocessing.SceneAnalysis{
				Timestamp:  scene.Timestamp,
				SceneType:  sanitizeUTF8(scene.SceneType),
				Description: sanitizeUTF8(scene.Description),
				Confidence: scene.Confidence,
			}
		}
	}

	// 转换物体检测
	if aiResult.Objects != nil {
		result.Objects = make([]*videoprocessing.ObjectDetection, len(aiResult.Objects))
		for i, obj := range aiResult.Objects {
			result.Objects[i] = &videoprocessing.ObjectDetection{
				Timestamp:  obj.Timestamp,
				ObjectName: sanitizeUTF8(obj.ObjectName),
				Confidence: obj.Confidence,
				Bbox: &videoprocessing.BoundingBox{
					X:      float64(obj.Bbox.X),
					Y:      float64(obj.Bbox.Y),
					Width:  float64(obj.Bbox.Width),
					Height: float64(obj.Bbox.Height),
				},
			}
		}
	}

	// 转换情感分析
	if aiResult.Emotions != nil {
		result.Emotions = make([]*videoprocessing.EmotionAnalysis, len(aiResult.Emotions))
		for i, emotion := range aiResult.Emotions {
			result.Emotions[i] = &videoprocessing.EmotionAnalysis{
				Timestamp:  emotion.Timestamp,
				Emotion:    sanitizeUTF8(emotion.Emotion),
				Confidence: emotion.Confidence,
				Description: sanitizeUTF8("检测到" + emotion.Emotion + "情感"),
			}
		}
	}

	// 转换文字识别
	if aiResult.Texts != nil {
		result.Texts = make([]*videoprocessing.TextRecognition, len(aiResult.Texts))
		for i, text := range aiResult.Texts {
			result.Texts[i] = &videoprocessing.TextRecognition{
				Timestamp: text.Timestamp,
				Text:      sanitizeUTF8(text.Text),
				Language:  text.Language,
				Confidence: text.Confidence,
				Bbox: &videoprocessing.BoundingBox{
					X:      float64(text.Bbox.X),
					Y:      float64(text.Bbox.Y),
					Width:  float64(text.Bbox.Width),
					Height: float64(text.Bbox.Height),
				},
			}
		}
	}

	// 转换音频分析
	if aiResult.Audio != nil {
		result.Audio = make([]*videoprocessing.AudioAnalysis, len(aiResult.Audio))
		for i, audio := range aiResult.Audio {
			result.Audio[i] = &videoprocessing.AudioAnalysis{
				Timestamp:    audio.Timestamp,
				Transcription: sanitizeUTF8(audio.Transcription),
				Language:     audio.Language,
				Confidence:   audio.Confidence,
			}
		}
	}

	// 转换视频总结
	if aiResult.Summary != nil {
		result.Summary = &videoprocessing.VideoSummary{
			Title:       sanitizeUTF8(aiResult.Summary.Title),
			Description: sanitizeUTF8(aiResult.Summary.Description),
			Keywords:    sanitizeUTF8Slice(aiResult.Summary.Keywords),
			Category:    aiResult.Summary.Category,
			Duration:    aiResult.Summary.Duration,
		}
	}

	return result, nil
}

// processVideoAnalysis 保留旧方法名作为兼容
func (l *AnalyzeVideoLogic) processVideoAnalysis(in *videoprocessing.AnalyzeVideoReq) (*videoprocessing.VideoAnalysisResult, error) {
	return l.processVideoAnalysisWithAI(in)
}

// getDefaultVideoAnalysisResult 返回默认的视频分析结果
func (l *AnalyzeVideoLogic) getDefaultVideoAnalysisResult(in *videoprocessing.AnalyzeVideoReq) *videoprocessing.AnalyzeVideoResp {
	return &videoprocessing.AnalyzeVideoResp{
		Status: 200,
		Msg:    "视频分析成功（使用模拟响应）",
		Result: &videoprocessing.VideoAnalysisResult{
			Scenes: []*videoprocessing.SceneAnalysis{
				{
					Timestamp:  0.0,
					SceneType:  "general",
					Confidence: 0.85,
					Description: sanitizeUTF8("视频内容分析中，由于AI服务暂时不可用，显示默认分析结果"),
				},
			},
			Objects:  []*videoprocessing.ObjectDetection{},
			Emotions: []*videoprocessing.EmotionAnalysis{},
			Texts:    []*videoprocessing.TextRecognition{},
			Audio:    []*videoprocessing.AudioAnalysis{},
			Summary: &videoprocessing.VideoSummary{
				Title:       sanitizeUTF8("视频分析结果"),
				Description: sanitizeUTF8("这是模拟的视频分析结果，实际环境中将调用AI服务进行详细分析"),
				Keywords:    []string{"分析", "视频", "AI"},
				Category:    "analysis",
				Duration:    60.0,
			},
		},
	}
}
