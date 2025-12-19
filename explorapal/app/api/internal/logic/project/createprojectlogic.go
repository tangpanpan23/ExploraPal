package project

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/model/hps"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建探索项目
func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProjectLogic) CreateProject(req *types.CreateProjectReq) (resp *types.CreateProjectResp, err error) {
	// 生成项目ID（简单实现，实际应该用雪花算法或其他分布式ID生成器）
	projectID := time.Now().UnixNano() / 1000000 // 毫秒级时间戳作为ID

	// 生成项目编码
	projectCode := fmt.Sprintf("P%d%03d", time.Now().Year(), projectID%1000)

	// 创建项目对象（使用生成的Projects结构体）
	project := &hps.Projects{
		ProjectId:   projectID,
		ProjectCode: projectCode,
		UserId:      req.UserId, // 注意：类型定义中使用的是UserId
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Category:    req.Category,
		Status:      "active", // 默认状态为活跃
		Progress:    0,        // 初始进度为0
		LastActivityAt: sql.NullTime{Valid: false}, // 初始时没有活动时间
	}

	// 设置标签
	if len(req.Tags) > 0 {
		tagsData, err := json.Marshal(req.Tags)
		if err != nil {
			l.Errorf("序列化项目标签失败: %v", err)
			return nil, err
		}
		project.Tags = sql.NullString{String: string(tagsData), Valid: true}
	}

	// 插入数据库
	_, err = l.svcCtx.ProjectModel.Insert(l.ctx, project)
	if err != nil {
		l.Errorf("创建项目失败: %v", err)
		return nil, err
	}

	l.Infof("项目创建成功: ID=%d, Code=%s, Title=%s", projectID, projectCode, req.Title)

	// 返回响应
	resp = &types.CreateProjectResp{
		ProjectId:   projectID, // 注意：响应中使用的是ProjectId
		ProjectCode: projectCode,
		Status:      "active",
	}

	return resp, nil
}
