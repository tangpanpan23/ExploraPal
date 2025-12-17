// 此文件为占位符，请运行 generate_proto.sh 生成实际的protobuf代码
// 生成protobuf代码后，此文件将被自动生成的代码替换

package aidialogue

import (
	"context"
	"google.golang.org/grpc"
)

// 占位符类型定义 - 这些将在运行protoc后自动生成
type AnalyzeImageReq struct {
	ImageUrl string
	Prompt   string
	Category string
}
type AnalyzeImageResp struct {
	Status        int32
	Msg           string
	ObjectName    string
	Category      string
	Confidence    float32
	Description   string
	KeyFeatures   []string
	ScientificName string
}

type GenerateQuestionsReq struct {
	ContextInfo string
	Category    string
	UserAge     int64
}
type GenerateQuestionsResp struct {
	Status    int32
	Msg       string
	Questions []*Question
}
type Question struct {
	Content    string
	Type       string
	Difficulty string
	Purpose    string
}

type PolishNoteReq struct {
	RawContent  string
	ContextInfo string
	Category    string
	UserAge     int64
}
type PolishNoteResp struct {
	Status        int32
	Msg           string
	Title         string
	Summary       string
	KeyPoints     []string
	FormattedText string
	Suggestions   []string
}

type GenerateReportReq struct {
	ProjectData string
	Category    string
}
type GenerateReportResp struct {
	Status     int32
	Msg        string
	Title      string
	Content    string
	Abstract   string
	Conclusion string
	NextSteps  []string
}

// 占位符服务接口
type AIDialogueServiceServer interface {
	AnalyzeImage(context.Context, *AnalyzeImageReq) (*AnalyzeImageResp, error)
	GenerateQuestions(context.Context, *GenerateQuestionsReq) (*GenerateQuestionsResp, error)
	PolishNote(context.Context, *PolishNoteReq) (*PolishNoteResp, error)
	GenerateReport(context.Context, *GenerateReportReq) (*GenerateReportResp, error)
}

type UnimplementedAIDialogueServiceServer struct{}

func (UnimplementedAIDialogueServiceServer) AnalyzeImage(context.Context, *AnalyzeImageReq) (*AnalyzeImageResp, error) {
	return nil, nil
}
func (UnimplementedAIDialogueServiceServer) GenerateQuestions(context.Context, *GenerateQuestionsReq) (*GenerateQuestionsResp, error) {
	return nil, nil
}
func (UnimplementedAIDialogueServiceServer) PolishNote(context.Context, *PolishNoteReq) (*PolishNoteResp, error) {
	return nil, nil
}
func (UnimplementedAIDialogueServiceServer) GenerateReport(context.Context, *GenerateReportReq) (*GenerateReportResp, error) {
	return nil, nil
}

// 占位符注册函数
func RegisterAIDialogueServiceServer(s *grpc.Server, srv AIDialogueServiceServer) {
	// 占位符实现，实际代码将由protoc生成
}
