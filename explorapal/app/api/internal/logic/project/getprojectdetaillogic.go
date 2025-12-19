package project

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取项目详情
func NewGetProjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectDetailLogic {
	return &GetProjectDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectDetailLogic) GetProjectDetail(req *types.GetProjectDetailReq) (resp *types.GetProjectDetailResp, err error) {
	// 获取项目详情
	project, err := l.svcCtx.ProjectModel.FindOneByProjectId(l.ctx, req.ProjectId)
	if err != nil {
		l.Errorf("获取项目详情失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp = &types.GetProjectDetailResp{
		Project: types.ProjectDetail{
			ProjectId:   project.ProjectId,
			ProjectCode: project.ProjectCode,
			Title:       project.Title,
			Description: project.Description.String, // sql.NullString
			Category:    project.Category,           // string 类型
			Status:      project.Status,
			Progress:    int32(project.Progress), // int64 -> int32
			CreateTime:  project.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  project.UpdateTime.Format("2006-01-02 15:04:05"),
			Tags:        []string{}, // TODO: 解析JSON标签
		},
		Activities:   []types.ProjectActivity{},   // TODO: 获取项目活动
		Achievements: []types.ProjectAchievement{}, // TODO: 获取项目成果
	}

	l.Infof("获取项目详情成功: ID=%d, Title=%s", req.ProjectId, project.Title)

	return resp, nil
}
