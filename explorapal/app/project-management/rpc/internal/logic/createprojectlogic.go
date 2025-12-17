package logic

import (
	"context"
	"fmt"
	"time"

	"explorapal/app/model/hps"
	"explorapal/app/project-management/rpc/internal/svc"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CreateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProjectLogic) CreateProject(in *projectmanagement.CreateProjectReq) (*projectmanagement.CreateProjectResp, error) {
	// 生成项目ID和项目编码
	projectID := time.Now().UnixNano()
	projectCode := l.generateProjectCode(projectID)

	// 创建项目记录
	project := &hps.Projects{
		ProjectID:   projectID,
		ProjectCode: projectCode,
		UserID:      in.UserId,
		Title:       in.Title,
		Description: in.Description,
		Category:    in.Category,
		Status:      "active",
		Progress:    0,
	}

	// 设置标签
	if len(in.Tags) > 0 {
		project.SetTags(in.Tags)
	}

	// 插入数据库
	_, err := l.svcCtx.ProjectModel.Insert(l.ctx, project)
	if err != nil {
		l.Logger.Errorf("创建项目失败: %v", err)
		return &projectmanagement.CreateProjectResp{
			Status: 500,
			Msg:    "创建项目失败",
		}, err
	}

	// 记录活动
	activity := &hps.ProjectActivities{
		ActivityID:  time.Now().UnixNano(),
		ProjectID:   projectID,
		UserID:      in.UserId,
		Type:        "create_project",
		Description: fmt.Sprintf("创建了项目：%s", in.Title),
	}
	_, err = l.svcCtx.ProjectActivityModel.Insert(l.ctx, activity)
	if err != nil {
		l.Logger.Errorf("记录项目活动失败: %v", err)
		// 不影响主要流程，只记录错误
	}

	l.Logger.Infof("用户 %d 创建项目成功: %s", in.UserId, projectCode)

	return &projectmanagement.CreateProjectResp{
		Status:      200,
		Msg:         "创建项目成功",
		ProjectId:   projectID,
		ProjectCode: projectCode,
	}, nil
}

// 生成项目编码
func (l *CreateProjectLogic) generateProjectCode(projectID int64) string {
	timestamp := time.Now().Format("20060102")
	return fmt.Sprintf("EXP-%s-%d", timestamp, projectID%100000)
}
