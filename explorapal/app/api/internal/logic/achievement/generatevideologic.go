package achievement

import (
	"context"
	"encoding/base64"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	videoprocessing "explorapal/app/video-processing/rpc"

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

func (l *GenerateVideoLogic) GenerateVideo(req *types.GenerateVideoReq) (resp *types.GenerateVideoResp, err error) {
	// TODO: 实现视频生成逻辑

	// 1. 调用视频处理RPC服务
	rpcReq := &videoprocessing.GenerateVideoReq{
		Script:  req.Script,
		Style:   req.Style,
		Duration: req.Duration,
		Scenes:  req.Scenes,
		Voice:   req.Voice,
		Language: req.Language,
	}

	rpcResp, err := l.svcCtx.VideoProcessingRpc.GenerateVideo(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("调用视频处理RPC服务失败: %v", err)
		// 返回模拟结果
		return l.getDefaultGenerateVideoResponse(req), nil
	}

	// 2. 转换响应格式
	videoDataBase64 := base64.StdEncoding.EncodeToString(rpcResp.VideoData)
	metadata := l.convertVideoMetadata(rpcResp.Metadata)

	return &types.GenerateVideoResp{
		VideoData:     videoDataBase64,
		Format:        rpcResp.Format,
		Duration:      rpcResp.Duration,
		Metadata:      *metadata,
		AchievementId: 0, // TODO: 保存到数据库
	}, nil
}

// convertVideoMetadata 转换视频元数据格式
func (l *GenerateVideoLogic) convertVideoMetadata(rpcMetadata *videoprocessing.VideoMetadata) *types.VideoMetadata {
	if rpcMetadata == nil {
		return &types.VideoMetadata{}
	}

	return &types.VideoMetadata{
		Title:          rpcMetadata.Title,
		Description:    rpcMetadata.Description,
		Scenes:         rpcMetadata.Scenes,
		AudioLanguage:  rpcMetadata.AudioLanguage,
		Resolution:     rpcMetadata.Resolution,
	}
}

// getDefaultGenerateVideoResponse 返回默认的视频生成响应
func (l *GenerateVideoLogic) getDefaultGenerateVideoResponse(req *types.GenerateVideoReq) *types.GenerateVideoResp {
	// 模拟视频数据
	mockVideoData := []byte("这是模拟的视频数据。实际环境中，这里将调用AI视频生成服务来生成真实的视频内容。")
	videoDataBase64 := base64.StdEncoding.EncodeToString(mockVideoData)

	return &types.GenerateVideoResp{
		VideoData:     videoDataBase64,
		Format:        "mp4",
		Duration:      req.Duration,
		Metadata: types.VideoMetadata{
			Title:          "AI生成的演示视频",
			Description:    "这是AI生成的演示视频，由于AI服务暂时不可用，显示默认内容",
			Scenes:         req.Scenes,
			AudioLanguage:  req.Language,
			Resolution:     "1920x1080",
		},
		AchievementId: 0,
	}
}