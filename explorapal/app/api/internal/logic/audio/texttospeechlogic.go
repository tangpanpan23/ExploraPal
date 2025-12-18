package audio

import (
	"context"
	"encoding/base64"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	audioprocessing "explorapal/app/audio-processing/proto"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type TextToSpeechLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTextToSpeechLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TextToSpeechLogic {
	return &TextToSpeechLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TextToSpeechLogic) TextToSpeech(req *types.TextToSpeechReq) (resp *types.TextToSpeechResp, err error) {
	// TODO: 实现文字转语音逻辑

	// 1. 创建音频处理RPC客户端
	audioRpcClient, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9004"}, // 音频处理RPC服务地址
		Timeout:   75000,                      // 75秒超时
	})
	if err != nil {
		l.Logger.Errorf("创建语音处理RPC客户端失败: %v", err)
		return l.getDefaultTextToSpeechResponse(req), nil
	}

	// 2. 准备RPC请求
	audioProcessingClient := audioprocessing.NewAudioProcessingServiceClient(audioRpcClient.Conn())
	rpcReq := &audioprocessing.TextToSpeechReq{
		Text:     req.Text,
		Voice:    req.Voice,
		Language: req.Language,
		Speed:    req.Speed,
	}

	// 3. 调用音频处理RPC服务
	rpcResp, err := audioProcessingClient.TextToSpeech(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("调用语音处理RPC服务失败: %v", err)
		// 返回模拟响应
		return l.getDefaultTextToSpeechResponse(req), nil
	}

	// 4. 将音频数据编码为base64
	audioDataBase64 := base64.StdEncoding.EncodeToString(rpcResp.AudioData)

	return &types.TextToSpeechResp{
		AudioData: audioDataBase64,
		Format:    rpcResp.Format,
		Duration:  0, // TODO: 计算音频时长
		ExpressionId: 0, // TODO: 保存到数据库
	}, nil
}

// getDefaultTextToSpeechResponse 返回默认的文字转语音响应
func (l *TextToSpeechLogic) getDefaultTextToSpeechResponse(req *types.TextToSpeechReq) *types.TextToSpeechResp {
	// 模拟音频数据（实际应该是真实的音频数据）
	mockAudioData := []byte("这是模拟的音频数据。实际环境中，这里将调用语音合成服务来生成真实的音频。")
	audioDataBase64 := base64.StdEncoding.EncodeToString(mockAudioData)

	return &types.TextToSpeechResp{
		AudioData: audioDataBase64,
		Format:    "wav",
		Duration:  float64(len(req.Text)) * 0.1, // 粗略估算时长
		ExpressionId: 0,
	}
}