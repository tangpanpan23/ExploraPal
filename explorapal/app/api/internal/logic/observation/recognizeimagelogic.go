package observation

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecognizeImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 识别图片内容
func NewRecognizeImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecognizeImageLogic {
	return &RecognizeImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecognizeImageLogic) RecognizeImage(req *types.RecognizeImageReq) (resp *types.RecognizeImageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
