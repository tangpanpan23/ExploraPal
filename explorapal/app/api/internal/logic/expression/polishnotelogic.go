package expression

import (
	"context"
	"encoding/json"
	"fmt"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PolishNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI润色生成笔记
func NewPolishNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PolishNoteLogic {
	return &PolishNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PolishNoteLogic) PolishNote(req *types.PolishNoteReq) (resp *types.PolishNoteResp, err error) {
	// 将context_info转换为字符串
	contextInfoStr := ""
	if req.ContextInfo != nil {
		contextBytes, err := json.Marshal(req.ContextInfo)
		if err != nil {
			l.Errorf("序列化context_info失败: %v", err)
			contextInfoStr = fmt.Sprintf("observation_results: %v, previous_answers: %v, project_category: %v",
				req.ContextInfo["observation_results"],
				req.ContextInfo["previous_answers"],
				req.ContextInfo["project_category"])
		} else {
			contextInfoStr = string(contextBytes)
		}
	}

	// 记录API层参数
	l.Infof("调用AI润色笔记: RawContent长度=%d, ContextInfo='%s'", len(req.RawContent), contextInfoStr)

	// 调用AI服务润色笔记
	aiResult, err := l.svcCtx.AIClient.PolishNote(l.ctx, req.RawContent, contextInfoStr)
	if err != nil {
		l.Errorf("AI润色笔记失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp = &types.PolishNoteResp{
		Title:     aiResult.Title,
		Summary:   aiResult.Summary,
		KeyPoints: aiResult.KeyPoints,
		FormattedText: fmt.Sprintf("%s\n\n%s", aiResult.Title, aiResult.Summary),
		Suggestions:   []string{"可以画一幅相关的插图", "试着讲给别人听"},
	}

	l.Infof("笔记润色完成，原始长度: %d, 润色后标题: %s", len(req.RawContent), aiResult.Title)

	return resp, nil
}
