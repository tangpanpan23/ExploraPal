package achievement

import (
	"context"
	"encoding/base64"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	videoprocessing "explorapal/app/video-processing/proto"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
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
	l.Logger.Infof("开始生成视频 - 项目ID: %d, 用户ID: %d", req.ProjectId, req.UserId)

	// 1. 创建视频处理RPC客户端
	videoRpcClient, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9005"}, // 视频处理RPC服务地址
		Timeout:   300000,                     // 300秒超时 (5分钟)
	})
	if err != nil {
		l.Logger.Errorf("创建视频处理RPC客户端失败: %v", err)
		return l.getDefaultGenerateVideoResponse(req), nil
	}

	// 2. 构建RPC请求
	rpcReq := &videoprocessing.GenerateVideoReq{
		Style:    req.Style,
		Duration: req.Duration,
		Scenes:   req.Scenes,
		Voice:    req.Voice,
		Language: req.Language,
	}

	// 根据输入模式设置不同的参数
	if req.ImageData != "" && req.Prompt != "" {
		// 图像到视频模式 - 使用豆包Doubao-Seedance-1.0-lite-i2v
		l.Logger.Infof("使用图像到视频模式 - 豆包Doubao-Seedance-1.0-lite-i2v")
		rpcReq.ImageData = req.ImageData
		rpcReq.Prompt = req.Prompt
		rpcReq.Script = "" // 图像模式下不需要脚本
	} else {
		// 文本到视频模式 - 使用原有逻辑
		l.Logger.Infof("使用文本到视频模式")
		rpcReq.Script = req.Script
		rpcReq.ImageData = ""
		rpcReq.Prompt = ""
	}

	// 3. 调用视频处理RPC服务
	videoProcessingClient := videoprocessing.NewVideoProcessingServiceClient(videoRpcClient.Conn())
	rpcResp, err := videoProcessingClient.GenerateVideo(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("调用视频处理RPC服务失败: %v", err)
		// 返回模拟结果
		return l.getDefaultGenerateVideoResponse(req), nil
	}

	// 4. 转换响应格式
	videoDataBase64 := base64.StdEncoding.EncodeToString(rpcResp.VideoData)
	metadata := l.convertVideoMetadata(rpcResp.Metadata)

	l.Logger.Infof("视频生成成功 - 格式: %s, 大小: %d bytes, 时长: %.2f秒",
		rpcResp.Format, len(rpcResp.VideoData), rpcResp.Duration)

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