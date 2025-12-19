package expression

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/audio-processing/proto"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type SpeechToTextLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 语音转文字
func NewSpeechToTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SpeechToTextLogic {
	return &SpeechToTextLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SpeechToTextLogic) SpeechToText(req *types.SpeechToTextReq) (resp *types.SpeechToTextResp, err error) {
	// 创建语音处理RPC客户端
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9004"}, // 语音处理服务地址
		Timeout:   10000,                       // 10秒超时
	})
	client := proto.NewAudioProcessingServiceClient(conn.Conn())

	// 调用语音转文字RPC
	rpcReq := &proto.SpeechToTextReq{
		AudioData: req.AudioData,
		Format:    req.AudioFormat,
		Language:  req.Language,
	}

	rpcResp, err := client.SpeechToText(l.ctx, rpcReq)
	if err != nil {
		l.Errorf("语音转文字RPC调用失败: %v", err)
		return nil, err
	}

	// 转换响应
	resp = &types.SpeechToTextResp{
		Text:       rpcResp.Text,
		Confidence: float64(rpcResp.Confidence),
		Language:   req.Language,
		Duration:   0, // TODO: 从音频数据计算时长
	}

	l.Infof("语音转文字完成: 语言=%s, 置信度=%.2f", req.Language, rpcResp.Confidence)

	return resp, nil
}
