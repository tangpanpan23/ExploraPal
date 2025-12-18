package questioning

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateQuestionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成引导问题
func NewGenerateQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateQuestionsLogic {
	return &GenerateQuestionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateQuestionsLogic) GenerateQuestions(req *types.GenerateQuestionsReq) (resp *types.GenerateQuestionsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
