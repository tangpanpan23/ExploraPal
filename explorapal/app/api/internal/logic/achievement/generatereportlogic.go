package achievement

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成研究简报
func NewGenerateReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateReportLogic {
	return &GenerateReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateReportLogic) GenerateReport(req *types.GenerateReportReq) (resp *types.GenerateReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
