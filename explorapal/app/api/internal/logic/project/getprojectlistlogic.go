package project

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/model/hps"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取项目列表
func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectListLogic) GetProjectList(req *types.GetProjectListReq) (resp *types.GetProjectListResp, err error) {
	// 设置默认分页参数
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}

	// 获取项目列表
	var projects []*hps.Projects
	var total int64

	if req.Category != "" {
		// 按类别查询
		projects, err = l.svcCtx.ProjectModel.FindByCategory(l.ctx, req.UserId, req.Category, page, pageSize)
	} else if req.Status != "" {
		// TODO: 添加按状态查询的方法
		projects, err = l.svcCtx.ProjectModel.FindByUserID(l.ctx, req.UserId, page, pageSize)
	} else {
		// 默认查询用户的所有项目
		projects, err = l.svcCtx.ProjectModel.FindByUserID(l.ctx, req.UserId, page, pageSize)
	}

	if err != nil {
		l.Errorf("获取项目列表失败: %v", err)
		return nil, err
	}

	// 转换项目列表
	projectList := make([]types.ProjectInfo, len(projects))
	for i, project := range projects {
		projectList[i] = types.ProjectInfo{
			ProjectId:   project.ProjectId,
			ProjectCode: project.ProjectCode,
			Title:       project.Title,
			Description: project.Description.String,
			Category:    project.Category.String,
			Status:      project.Status,
			Progress:    int32(project.Progress),
			CreateTime:  project.CreateTime.Unix(),
			UpdateTime:  project.UpdateTime.Unix(),
		}
	}

	resp = &types.GetProjectListResp{
		Projects: projectList,
		Total:    int32(total), // TODO: 获取总数
		Page:     int32(page),
		PageSize: int32(pageSize),
	}

	l.Infof("获取项目列表成功: 用户ID=%d, 返回%d个项目", req.UserId, len(projectList))

	return resp, nil
}
