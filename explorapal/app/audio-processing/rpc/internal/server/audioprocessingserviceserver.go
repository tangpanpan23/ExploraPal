package server

import (
	"context"

	"explorapal/app/audio-processing/proto"
	"explorapal/app/audio-processing/rpc/internal/svc"

	logic "explorapal/app/audio-processing/rpc/internal/logic"
)

type AudioProcessingServiceServer struct {
	svcCtx *svc.ServiceContext
	audioprocessing.UnimplementedAudioProcessingServiceServer
}

func NewAudioProcessingServiceServer(svcCtx *svc.ServiceContext) *AudioProcessingServiceServer {
	return &AudioProcessingServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *AudioProcessingServiceServer) SpeechToText(ctx context.Context, req *audioprocessing.SpeechToTextReq) (*audioprocessing.SpeechToTextResp, error) {
	l := logic.NewSpeechToTextLogic(ctx, s.svcCtx)
	return l.SpeechToText(req)
}

func (s *AudioProcessingServiceServer) TextToSpeech(ctx context.Context, req *audioprocessing.TextToSpeechReq) (*audioprocessing.TextToSpeechResp, error) {
	l := logic.NewTextToSpeechLogic(ctx, s.svcCtx)
	return l.TextToSpeech(req)
}
