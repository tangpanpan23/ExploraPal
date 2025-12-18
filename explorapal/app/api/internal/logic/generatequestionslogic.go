package logic

import (
	"context"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type GenerateQuestionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateQuestionsLogic {
	return &GenerateQuestionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateQuestionsLogic) GenerateQuestions(req *types.GenerateQuestionsReq) (resp *types.GenerateQuestionsResp, err error) {
	// 调用AI对话RPC服务生成问题
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9002"}, // AI对话RPC服务地址
		Timeout:   75000,                      // 75秒超时，与AI分析超时保持一致
	})
	if err != nil {
		l.Logger.Errorf("创建AI RPC客户端失败: %v", err)
		return nil, err
	}

	aiClient := aidialogue.NewAIDialogueServiceClient(client.Conn())

	rpcResp, err := aiClient.GenerateQuestions(l.ctx, &aidialogue.GenerateQuestionsReq{
		ContextInfo: req.ContextInfo,
		Category:    req.Category,
		UserAge:     int64(req.UserAge),
	})
	if err != nil {
		l.Logger.Errorf("调用AI生成问题服务失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	questions := make([]types.Question, 0, len(rpcResp.Questions))
	for _, item := range rpcResp.Questions {
		questions = append(questions, types.Question{
			Content:    item.Content,
			Type:       item.Type,
			Difficulty: item.Difficulty,
			Purpose:    item.Purpose,
		})
	}

	resp = &types.GenerateQuestionsResp{
		Questions: questions,
	}

	return resp, nil
}
