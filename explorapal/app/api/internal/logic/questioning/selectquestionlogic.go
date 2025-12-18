package questioning

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 选择问题并获取AI回答
func NewSelectQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectQuestionLogic {
	return &SelectQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectQuestionLogic) SelectQuestion(req *types.SelectQuestionReq) (resp *types.SelectQuestionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
