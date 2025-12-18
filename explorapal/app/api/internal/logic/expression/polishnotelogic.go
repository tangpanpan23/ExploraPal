package expression

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
