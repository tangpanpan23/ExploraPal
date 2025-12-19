package expression

import (
	"context"
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
	// 调用AI服务润色笔记
	aiResult, err := l.svcCtx.AIClient.PolishNote(l.ctx, req.RawContent, req.ContextInfo)
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
