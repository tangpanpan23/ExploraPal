package project

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新项目状态
func NewUpdateProjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectStatusLogic {
	return &UpdateProjectStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProjectStatusLogic) UpdateProjectStatus(req *types.UpdateProjectStatusReq) (resp *types.CommonStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
