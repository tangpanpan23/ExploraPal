package observation

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadObservationImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传观察图片
func NewUploadObservationImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadObservationImageLogic {
	return &UploadObservationImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadObservationImageLogic) UploadObservationImage(req *types.UploadObservationImageReq) (resp *types.UploadObservationImageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
