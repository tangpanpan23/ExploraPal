package audio

import (
	"context"
	"encoding/base64"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/audio-processing/rpc"

	"github.com/zeromicro/go-zero/core/logx"
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

	// 1. 准备RPC请求
	rpcReq := &audioprocessing.TextToSpeechReq{
		Text:     req.Text,
		Voice:    req.Voice,
		Language: req.Language,
		Speed:    req.Speed,
	}

	// 2. 调用音频处理RPC服务
	rpcResp, err := l.svcCtx.AudioProcessingRpc.TextToSpeech(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("调用语音处理RPC服务失败: %v", err)
		// 返回模拟响应
		return l.getDefaultTextToSpeechResponse(req), nil
	}

	// 3. 将音频数据编码为base64
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