package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"explorapal/app/model/hps"
	"explorapal/app/project-management/rpc/internal/svc"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
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
		ProjectId:   projectID,
		ProjectCode: projectCode,
		UserId:      in.UserId,
		Title:       in.Title,
		Category:    in.Category,
		Status:      "active",
		Progress:    0,
	}

	// 设置描述
	if in.Description != "" {
		project.Description = sql.NullString{String: in.Description, Valid: true}
	}

	// 设置标签
	if len(in.Tags) > 0 {
		tagsJSON, _ := json.Marshal(in.Tags)
		project.Tags = sql.NullString{String: string(tagsJSON), Valid: true}
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
		ActivityId:  time.Now().UnixNano(),
		ProjectId:   projectID,
		UserId:      in.UserId,
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
