package achievement

import (
	"context"
	"encoding/base64"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	videoprocessing "explorapal/app/video-processing/rpc"

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

func (l *AnalyzeVideoLogic) AnalyzeVideo(req *types.AnalyzeVideoReq) (resp *types.AnalyzeVideoResp, err error) {
	// TODO: 实现视频分析逻辑

	// 1. 解码base64视频数据
	videoData, err := base64.StdEncoding.DecodeString(req.VideoData)
	if err != nil {
		l.Logger.Errorf("解码视频数据失败: %v", err)
		return nil, err
	}

	// 2. 调用视频处理RPC服务
	rpcReq := &videoprocessing.AnalyzeVideoReq{
		VideoData:    videoData,
		Format:       req.VideoFormat,
		AnalysisType: req.AnalysisType,
		Duration:     req.Duration,
	}

	rpcResp, err := l.svcCtx.VideoProcessingRpc.AnalyzeVideo(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("调用视频处理RPC服务失败: %v", err)
		// 返回模拟结果
		return l.getDefaultAnalyzeVideoResponse(req), nil
	}

	// 3. 转换响应格式
	videoAnalysis := l.convertVideoAnalysisResult(rpcResp.Result)

	return &types.AnalyzeVideoResp{
		VideoAnalysis: *videoAnalysis,
		AchievementId: 0, // TODO: 保存到数据库
	}, nil
}

// convertVideoAnalysisResult 转换视频分析结果格式
func (l *AnalyzeVideoLogic) convertVideoAnalysisResult(rpcResult *videoprocessing.VideoAnalysisResult) *types.VideoAnalysisResult {
	if rpcResult == nil {
		return &types.VideoAnalysisResult{}
	}

	// 转换场景分析
	scenes := make([]types.SceneAnalysis, len(rpcResult.Scenes))
	for i, scene := range rpcResult.Scenes {
		scenes[i] = types.SceneAnalysis{
			Timestamp:  scene.Timestamp,
			SceneType:  scene.SceneType,
			Confidence: scene.Confidence,
			Description: scene.Description,
		}
	}

	// 转换物体检测
	objects := make([]types.ObjectDetection, len(rpcResult.Objects))
	for i, obj := range rpcResult.Objects {
		objects[i] = types.ObjectDetection{
			Timestamp:  obj.Timestamp,
			ObjectName: obj.ObjectName,
			Confidence: obj.Confidence,
			Bbox: types.BoundingBox{
				X:      obj.Bbox.X,
				Y:      obj.Bbox.Y,
				Width:  obj.Bbox.Width,
				Height: obj.Bbox.Height,
			},
		}
	}

	// 转换情感分析
	emotions := make([]types.EmotionAnalysis, len(rpcResult.Emotions))
	for i, emotion := range rpcResult.Emotions {
		emotions[i] = types.EmotionAnalysis{
			Timestamp:  emotion.Timestamp,
			Emotion:    emotion.Emotion,
			Confidence: emotion.Confidence,
			Description: emotion.Description,
		}
	}

	// 转换文字识别
	texts := make([]types.TextRecognition, len(rpcResult.Texts))
	for i, text := range rpcResult.Texts {
		texts[i] = types.TextRecognition{
			Timestamp:  text.Timestamp,
			Text:       text.Text,
			Language:   text.Language,
			Confidence: text.Confidence,
			Bbox: types.BoundingBox{
				X:      text.Bbox.X,
				Y:      text.Bbox.Y,
				Width:  text.Bbox.Width,
				Height: text.Bbox.Height,
			},
		}
	}

	// 转换音频分析
	audio := make([]types.AudioAnalysis, len(rpcResult.Audio))
	for i, audioItem := range rpcResult.Audio {
		audio[i] = types.AudioAnalysis{
			Timestamp:    audioItem.Timestamp,
			Transcription: audioItem.Transcription,
			Language:     audioItem.Language,
			Confidence:   audioItem.Confidence,
		}
	}

	// 转换视频总结
	summary := types.VideoSummary{
		Title:       rpcResult.Summary.Title,
		Description: rpcResult.Summary.Description,
		Keywords:    rpcResult.Summary.Keywords,
		Category:    rpcResult.Summary.Category,
		Duration:    rpcResult.Summary.Duration,
	}

	return &types.VideoAnalysisResult{
		Scenes:   scenes,
		Objects:  objects,
		Emotions: emotions,
		Texts:    texts,
		Audio:    audio,
		Summary:  summary,
	}
}

// getDefaultAnalyzeVideoResponse 返回默认的视频分析响应
func (l *AnalyzeVideoLogic) getDefaultAnalyzeVideoResponse(req *types.AnalyzeVideoReq) *types.AnalyzeVideoResp {
	return &types.AnalyzeVideoResp{
		VideoAnalysis: types.VideoAnalysisResult{
			Scenes: []types.SceneAnalysis{
				{
					Timestamp:  0.0,
					SceneType:  "educational",
					Confidence: 0.85,
					Description: "视频分析中，由于AI服务暂时不可用，显示默认分析结果",
				},
			},
			Objects:  []types.ObjectDetection{},
			Emotions: []types.EmotionAnalysis{},
			Texts:    []types.TextRecognition{},
			Audio:    []types.AudioAnalysis{},
			Summary: types.VideoSummary{
				Title:       "视频分析结果",
				Description: "这是模拟的视频分析结果，实际环境中将调用AI服务进行详细分析",
				Keywords:    []string{"分析", "视频", "AI"},
				Category:    "analysis",
				Duration:    60.0,
			},
		},
		AchievementId: 0,
	}
}