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
		l.Logger.Errorf("笔记润色失败: %v", err)
		return &aidialogue.PolishNoteResp{
			Status: 500,
			Msg:    "笔记润色失败",
		}, err
	}

	return &aidialogue.PolishNoteResp{
		Status:        200,
		Msg:           "笔记润色成功",
		Title:         result.Title,
		Summary:       result.Summary,
		KeyPoints:     result.KeyPoints,
		FormattedText: result.FormattedText,
		Suggestions:   []string{}, // TODO: 从result中提取建议
	}, nil
}
