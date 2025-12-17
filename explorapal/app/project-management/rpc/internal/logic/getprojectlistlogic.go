package logic

import (
	"context"
	"encoding/json"

	"explorapal/app/project-management/rpc/internal/svc"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectListLogic) GetProjectList(in *projectmanagement.GetProjectListReq) (*projectmanagement.GetProjectListResp, error) {
	var projects []*projectmanagement.ProjectInfo
	var total int64

	// 根据筛选条件获取项目列表
	if in.Category != "" {
		// 按类别筛选
		projectList, err := l.svcCtx.ProjectModel.FindByCategory(l.ctx, in.UserId, in.Category, in.Page, in.PageSize)
		if err != nil {
			l.Logger.Errorf("按类别获取项目列表失败: %v", err)
			return &projectmanagement.GetProjectListResp{
				Status: 500,
				Msg:    "获取项目列表失败",
			}, err
		}

		// 转换数据结构
		for _, p := range projectList {
			var tags []string
			if p.Tags.Valid {
				_ = json.Unmarshal([]byte(p.Tags.String), &tags)
			}

			projectInfo := &projectmanagement.ProjectInfo{
				ProjectId:   p.ProjectId,
				ProjectCode: p.ProjectCode,
				Title:       p.Title,
				Category:    p.Category,
				Status:      p.Status,
				Progress:    int32(p.Progress),
				CreateTime:  p.CreateTime.Format("2006-01-02 15:04:05"),
				UpdateTime:  p.UpdateTime.Format("2006-01-02 15:04:05"),
				Tags:        tags,
			}

			if p.Description.Valid {
				projectInfo.Description = p.Description.String
			}
			if p.LastActivityAt.Valid {
				projectInfo.LastActivity = p.LastActivityAt.Time.Format("2006-01-02 15:04:05")
			}

			projects = append(projects, projectInfo)
		}

		// 这里简化处理，实际应该查询总数
		total = int64(len(projects))
	} else {
		// 获取所有项目
		projectList, err := l.svcCtx.ProjectModel.FindByUserID(l.ctx, in.UserId, in.Page, in.PageSize)
		if err != nil {
			l.Logger.Errorf("获取项目列表失败: %v", err)
			return &projectmanagement.GetProjectListResp{
				Status: 500,
				Msg:    "获取项目列表失败",
			}, err
		}

		// 转换数据结构
		for _, p := range projectList {
			var tags []string
			if p.Tags.Valid {
				_ = json.Unmarshal([]byte(p.Tags.String), &tags)
			}

			projectInfo := &projectmanagement.ProjectInfo{
				ProjectId:   p.ProjectId,
				ProjectCode: p.ProjectCode,
				Title:       p.Title,
				Category:    p.Category,
				Status:      p.Status,
				Progress:    int32(p.Progress),
				CreateTime:  p.CreateTime.Format("2006-01-02 15:04:05"),
				UpdateTime:  p.UpdateTime.Format("2006-01-02 15:04:05"),
				Tags:        tags,
			}

			if p.Description.Valid {
				projectInfo.Description = p.Description.String
			}
			if p.LastActivityAt.Valid {
				projectInfo.LastActivity = p.LastActivityAt.Time.Format("2006-01-02 15:04:05")
			}

			projects = append(projects, projectInfo)
		}

		// 这里简化处理，实际应该查询总数
		total = int64(len(projects))
	}

	return &projectmanagement.GetProjectListResp{
		Status:   200,
		Msg:      "获取项目列表成功",
		List:     projects,
		Total:    total,
		PageSize: in.PageSize,
		Page:     in.Page,
	}, nil
}
