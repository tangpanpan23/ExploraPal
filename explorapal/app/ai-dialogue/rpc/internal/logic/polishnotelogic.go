package logic

import (
	"context"

	"explorapal/app/ai-dialogue/rpc/aidialogue"
	"explorapal/app/ai-dialogue/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *PolishNoteLogic) PolishNote(in *aidialogue.PolishNoteReq) (*aidialogue.PolishNoteResp, error) {
	// TODO: 实现笔记润色逻辑
	result, err := l.svcCtx.AIClient.PolishNote(l.ctx, in.RawContent, in.ContextInfo)
	if err != nil {
		l.Logger.Errorf("笔记润色失败，使用默认润色: %v", err)
		// 返回默认的润色结果
		return l.getDefaultPolishedNoteResponse(in.RawContent, in.ContextInfo), nil
	}

	return &aidialogue.PolishNoteResp{
		Status:        200,
		Msg:           "笔记润色成功",
		Title:         sanitizeUTF8(result.Title),
		Summary:       sanitizeUTF8(result.Summary),
		KeyPoints:     sanitizeUTF8Slice(result.KeyPoints),
		FormattedText: sanitizeUTF8(result.FormattedText),
		Suggestions:   []string{}, // TODO: 从result中提取建议
	}, nil
}

// getDefaultPolishedNoteResponse 返回默认的笔记润色结果
func (l *PolishNoteLogic) getDefaultPolishedNoteResponse(rawContent, contextInfo string) *aidialogue.PolishNoteResp {
	return &aidialogue.PolishNoteResp{
		Status:        200,
		Msg:           "笔记润色成功（使用模拟响应）",
		Title:         sanitizeUTF8("探索笔记"),
		Summary:       sanitizeUTF8("这是孩子记录的探索笔记，由于AI服务暂时不可用，显示原始内容。"),
		KeyPoints:     sanitizeUTF8Slice([]string{"记录观察", "表达想法", "提出问题"}),
		FormattedText: sanitizeUTF8(rawContent), // 保持原始内容
		Suggestions:   sanitizeUTF8Slice([]string{"可以添加更多观察细节", "可以画一幅相关的图画", "可以想一想为什么会这样"}),
	}
}

