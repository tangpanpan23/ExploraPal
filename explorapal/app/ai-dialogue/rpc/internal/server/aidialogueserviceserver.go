package server

import (
	"context"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/ai-dialogue/rpc/internal/logic"
	"explorapal/app/ai-dialogue/rpc/internal/svc"
)

type AIDialogueServiceServer struct {
	svcCtx *svc.ServiceContext
	aidialogue.UnimplementedAIDialogueServiceServer
}

func NewAIDialogueServiceServer(svcCtx *svc.ServiceContext) *AIDialogueServiceServer {
	return &AIDialogueServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *AIDialogueServiceServer) AnalyzeImage(ctx context.Context, in *aidialogue.AnalyzeImageReq) (*aidialogue.AnalyzeImageResp, error) {
	l := logic.NewAnalyzeImageLogic(ctx, s.svcCtx)
	return l.AnalyzeImage(in)
}

func (s *AIDialogueServiceServer) GenerateQuestions(ctx context.Context, in *aidialogue.GenerateQuestionsReq) (*aidialogue.GenerateQuestionsResp, error) {
	l := logic.NewGenerateQuestionsLogic(ctx, s.svcCtx)
	return l.GenerateQuestions(in)
}

func (s *AIDialogueServiceServer) PolishNote(ctx context.Context, in *aidialogue.PolishNoteReq) (*aidialogue.PolishNoteResp, error) {
	l := logic.NewPolishNoteLogic(ctx, s.svcCtx)
	return l.PolishNote(in)
}

func (s *AIDialogueServiceServer) GenerateReport(ctx context.Context, in *aidialogue.GenerateReportReq) (*aidialogue.GenerateReportResp, error) {
	l := logic.NewGenerateReportLogic(ctx, s.svcCtx)
	return l.GenerateReport(in)
}
