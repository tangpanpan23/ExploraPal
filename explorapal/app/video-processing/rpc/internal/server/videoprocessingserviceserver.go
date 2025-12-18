package server

import (
	"context"

	videoprocessing "explorapal/app/video-processing/proto"
	"explorapal/app/video-processing/rpc/internal/svc"

	logic "explorapal/app/video-processing/rpc/internal/logic"
)

type VideoProcessingServiceServer struct {
	svcCtx *svc.ServiceContext
	videoprocessing.UnimplementedVideoProcessingServiceServer
}

func NewVideoProcessingServiceServer(svcCtx *svc.ServiceContext) *VideoProcessingServiceServer {
	return &VideoProcessingServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoProcessingServiceServer) AnalyzeVideo(ctx context.Context, req *videoprocessing.AnalyzeVideoReq) (*videoprocessing.AnalyzeVideoResp, error) {
	l := logic.NewAnalyzeVideoLogic(ctx, s.svcCtx)
	return l.AnalyzeVideo(req)
}

func (s *VideoProcessingServiceServer) GenerateVideo(ctx context.Context, req *videoprocessing.GenerateVideoReq) (*videoprocessing.GenerateVideoResp, error) {
	l := logic.NewGenerateVideoLogic(ctx, s.svcCtx)
	return l.GenerateVideo(req)
}
