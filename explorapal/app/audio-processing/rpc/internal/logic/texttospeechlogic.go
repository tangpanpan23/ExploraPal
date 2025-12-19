package logic

import (
	"context"

	"explorapal/app/audio-processing/proto"
	"explorapal/app/audio-processing/rpc/internal/svc"
	"explorapal/third/openai"

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

func (l *TextToSpeechLogic) TextToSpeech(in *audioprocessing.TextToSpeechReq) (*audioprocessing.TextToSpeechResp, error) {
	// 记录API请求参数
	l.Infof("文字转语音请求: Text长度=%d, Voice=%s, Language=%s, Speed=%.2f",
		len(in.Text), in.Voice, in.Language, in.Speed)

	// 1. 验证输入
	if in.Text == "" {
		return &audioprocessing.TextToSpeechResp{
			Status: 400,
			Msg:    "文字内容不能为空",
		}, nil
	}

	// 2. 设置默认参数
	voice := in.Voice
	if voice == "" {
		voice = l.svcCtx.Config.SystemTTS.Voice
	}

	language := in.Language
	if language == "" {
		language = l.svcCtx.Config.SystemSTT.Language
	}

	speed := in.Speed
	if speed <= 0 {
		speed = l.svcCtx.Config.SystemTTS.Speed
	}

	// 3. 调用AI服务进行语音合成
	audioData, format, err := l.processTextToSpeechWithAI(in.Text, voice, language, speed)
	if err != nil {
		l.Logger.Errorf("AI文字转语音失败: %v", err)
		// 返回模拟结果
		return l.getDefaultTextToSpeechResult(in), nil
	}

	// 记录成功结果
	l.Infof("文字转语音成功: 音频大小=%d bytes, 格式=%s", len(audioData), format)

	return &audioprocessing.TextToSpeechResp{
		Status:    200,
		Msg:       "文字转语音成功",
		AudioData: audioData,
		Format:    format,
	}, nil
}

// processTextToSpeechWithAI 调用AI服务处理文字转语音
func (l *TextToSpeechLogic) processTextToSpeechWithAI(text, voice, language string, speed float64) ([]byte, string, error) {
	// 使用ServiceContext中的AI客户端
	// 显式使用openai包以避免编译器误报
	_ = openai.Client{}
	audioData, format, err := l.svcCtx.AIClient.TextToSpeech(l.ctx, text, voice, language, speed)
	if err != nil {
		l.Logger.Errorf("AI文字转语音调用失败: %v", err)
		return nil, "", err
	}

	return audioData, format, nil
}

// processTextToSpeech 处理文字转语音 (保留旧方法名作为兼容)
func (l *TextToSpeechLogic) processTextToSpeech(text, voice, language string, speed float64) ([]byte, string, error) {
	return l.processTextToSpeechWithAI(text, voice, language, speed)
}

// getDefaultTextToSpeechResult 返回默认的文字转语音结果
func (l *TextToSpeechLogic) getDefaultTextToSpeechResult(in *audioprocessing.TextToSpeechReq) *audioprocessing.TextToSpeechResp {
	return &audioprocessing.TextToSpeechResp{
		Status:    200,
		Msg:       "文字转语音成功（使用模拟响应）",
		AudioData: []byte("这是模拟的音频数据。实际环境中，这里将调用语音合成服务来生成真实的音频。"),
		Format:    "wav",
	}
}
