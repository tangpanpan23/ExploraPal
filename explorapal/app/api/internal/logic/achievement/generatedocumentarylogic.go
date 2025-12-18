package achievement

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDocumentaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成纪录片脚本
func NewGenerateDocumentaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDocumentaryLogic {
	return &GenerateDocumentaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateDocumentaryLogic) GenerateDocumentary(req *types.GenerateDocumentaryReq) (resp *types.GenerateDocumentaryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
