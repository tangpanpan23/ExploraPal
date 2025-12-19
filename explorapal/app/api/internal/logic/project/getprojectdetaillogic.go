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
		ProjectId:   project.ProjectId,
		ProjectCode: project.ProjectCode,
		UserId:      project.UserId,
		Title:       project.Title,
		Description: project.Description.String, // sql.NullString
		Category:    project.Category.String,     // sql.NullString
		Status:      project.Status,
		Progress:    int32(project.Progress), // int64 -> int32
		Tags:        []string{},              // TODO: 解析JSON标签
		CreateTime:  project.CreateTime.Unix(),
		UpdateTime:  project.UpdateTime.Unix(),
	}

	l.Infof("获取项目详情成功: ID=%d, Title=%s", req.ProjectId, project.Title)

	return resp, nil
}
