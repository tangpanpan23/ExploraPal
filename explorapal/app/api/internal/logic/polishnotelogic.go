package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type PolishNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPolishNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PolishNoteLogic {
	return &PolishNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PolishNoteLogic) PolishNote(req *types.PolishNoteReq) (resp *types.PolishNoteResp, err error) {
	// 将context_info转换为字符串
	contextInfoStr := ""
	if req.ContextInfo != nil {
		contextBytes, err := json.Marshal(req.ContextInfo)
		if err != nil {
			l.Logger.Errorf("序列化context_info失败: %v", err)
			contextInfoStr = fmt.Sprintf("observation_results: %v, previous_answers: %v, project_category: %v",
				req.ContextInfo["observation_results"],
				req.ContextInfo["previous_answers"],
				req.ContextInfo["project_category"])
		} else {
			contextInfoStr = string(contextBytes)
		}
	}

	// 调用AI对话RPC服务润色笔记
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9002"}, // AI对话RPC服务地址
		Timeout:   75000,                      // 75秒超时，与AI分析超时保持一致
	})
	if err != nil {
		l.Logger.Errorf("创建AI RPC客户端失败: %v", err)
		return nil, err
	}

	aiClient := aidialogue.NewAIDialogueServiceClient(client.Conn())

	rpcResp, err := aiClient.PolishNote(l.ctx, &aidialogue.PolishNoteReq{
		RawContent:  req.RawContent,
		ContextInfo: contextInfoStr, // 使用转换后的字符串
		Category:    req.Category,
		UserAge:     int64(req.UserAge),
	})
	if err != nil {
		l.Logger.Errorf("调用AI笔记润色服务失败: %v", err)
		return nil, err
	}

	resp = &types.PolishNoteResp{
		Title:         rpcResp.Title,
		Summary:       rpcResp.Summary,
		KeyPoints:     rpcResp.KeyPoints,
		FormattedText: rpcResp.FormattedText,
		Suggestions:   rpcResp.Suggestions,
	}

	return resp, nil
}
