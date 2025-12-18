package logic

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
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

func (l *GetProjectDetailLogic) GetProjectDetail(req *types.GetProjectDetailReq) (resp *types.GetProjectDetailResp, err error) {
	// 调用项目管理RPC服务
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9001"}, // 项目管理RPC服务地址
	})
	if err != nil {
		l.Logger.Errorf("创建RPC客户端失败: %v", err)
		return nil, err
	}

	projectClient := projectmanagement.NewProjectManagementServiceClient(client.Conn())

	rpcResp, err := projectClient.GetProjectDetail(l.ctx, &projectmanagement.GetProjectDetailReq{
		ProjectId: req.ProjectId,
		UserId:    req.UserId,
	})
	if err != nil {
		l.Logger.Errorf("调用项目管理RPC服务失败: %v", err)
		return nil, err
	}

	// 转换项目详情
	project := types.ProjectDetail{
		ProjectId:    rpcResp.Project.ProjectId,
		ProjectCode:  rpcResp.Project.ProjectCode,
		Title:        rpcResp.Project.Title,
		Description:  rpcResp.Project.Description,
		Category:     rpcResp.Project.Category,
		Status:       rpcResp.Project.Status,
		Progress:     rpcResp.Project.Progress,
		CreateTime:   rpcResp.Project.CreateTime,
		UpdateTime:   rpcResp.Project.UpdateTime,
		LastActivity: rpcResp.Project.LastActivity,
		Tags:         rpcResp.Project.Tags,
	}

	// 转换活动记录
	activities := make([]types.ProjectActivity, 0, len(rpcResp.Activities))
	for _, item := range rpcResp.Activities {
		activities = append(activities, types.ProjectActivity{
			ActivityId:  item.ActivityId,
			Type:        item.Type,
			Description: item.Description,
			CreateTime:  item.CreateTime,
		})
	}

	// 转换成果
	achievements := make([]types.ProjectAchievement, 0, len(rpcResp.Achievements))
	for _, item := range rpcResp.Achievements {
		achievements = append(achievements, types.ProjectAchievement{
			AchievementId: item.AchievementId,
			Type:          item.Type,
			Title:         item.Title,
			Content:       item.Content,
			Url:           item.Url,
			CreateTime:    item.CreateTime,
		})
	}

	resp = &types.GetProjectDetailResp{
		Project:      project,
		Activities:   activities,
		Achievements: achievements,
	}

	return resp, nil
}
