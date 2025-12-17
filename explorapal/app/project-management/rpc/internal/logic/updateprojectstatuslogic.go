package logic

import (
	"context"

	"explorapal/app/project-management/rpc/internal/svc"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectStatusLogic {
	return &UpdateProjectStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProjectStatusLogic) UpdateProjectStatus(in *projectmanagement.UpdateProjectStatusReq) (*projectmanagement.UpdateProjectStatusResp, error) {
	// 获取现有项目
	project, err := l.svcCtx.ProjectModel.FindOneByProjectId(l.ctx, in.ProjectId)
	if err != nil {
		l.Logger.Errorf("获取项目失败: %v", err)
		return &projectmanagement.UpdateProjectStatusResp{
			Status: 500,
			Msg:    "获取项目失败",
		}, err
	}

	// 更新状态
	project.Status = in.Status
	err = l.svcCtx.ProjectModel.Update(l.ctx, project)
	if err != nil {
		l.Logger.Errorf("更新项目状态失败: %v", err)
		return &projectmanagement.UpdateProjectStatusResp{
			Status: 500,
			Msg:    "更新项目状态失败",
		}, err
	}

	return &projectmanagement.UpdateProjectStatusResp{
		Status: 200,
		Msg:    "更新项目状态成功",
	}, nil
}
