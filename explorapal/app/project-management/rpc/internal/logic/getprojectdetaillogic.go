package logic

import (
	"context"
	"encoding/json"

	"explorapal/app/project-management/rpc/internal/svc"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectDetailLogic {
	return &GetProjectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectDetailLogic) GetProjectDetail(in *projectmanagement.GetProjectDetailReq) (*projectmanagement.GetProjectDetailResp, error) {
	// 获取项目基本信息
	project, err := l.svcCtx.ProjectModel.FindOneByProjectId(l.ctx, in.ProjectId)
	if err != nil {
		l.Logger.Errorf("获取项目详情失败: %v", err)
		return &projectmanagement.GetProjectDetailResp{
			Status: 500,
			Msg:    "获取项目详情失败",
		}, err
	}

	// 解析标签
	var tags []string
	if project.Tags.Valid {
		_ = json.Unmarshal([]byte(project.Tags.String), &tags)
	}

	// 转换项目信息
	projectDetail := &projectmanagement.ProjectDetail{
		ProjectId:   project.ProjectId,
		ProjectCode: project.ProjectCode,
		Title:       project.Title,
		Category:    project.Category,
		Status:      project.Status,
		Progress:    int32(project.Progress),
		CreateTime:  project.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  project.UpdateTime.Format("2006-01-02 15:04:05"),
		Tags:        tags,
	}

	// 设置描述
	if project.Description.Valid {
		projectDetail.Description = project.Description.String
	}

	// 设置最后活动时间
	if project.LastActivityAt.Valid {
		projectDetail.LastActivity = project.LastActivityAt.Time.Format("2006-01-02 15:04:05")
	}

	return &projectmanagement.GetProjectDetailResp{
		Status:  200,
		Msg:     "获取项目详情成功",
		Project: projectDetail,
	}, nil
}
