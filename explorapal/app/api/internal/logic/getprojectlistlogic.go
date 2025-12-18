package logic

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
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

func (l *GetProjectListLogic) GetProjectList(req *types.GetProjectListReq) (resp *types.GetProjectListResp, err error) {
	// 调用项目管理RPC服务
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9001"}, // 项目管理RPC服务地址
	})
	if err != nil {
		l.Logger.Errorf("创建RPC客户端失败: %v", err)
		return nil, err
	}

	projectClient := projectmanagement.NewProjectManagementServiceClient(client.Conn())

	rpcResp, err := projectClient.GetProjectList(l.ctx, &projectmanagement.GetProjectListReq{
		UserId:   req.UserId,
		Category: req.Category,
		Status:   req.Status,
		PageSize: req.PageSize,
		Page:     req.Page,
	})
	if err != nil {
		l.Logger.Errorf("调用项目管理RPC服务失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	projectList := make([]types.ProjectInfo, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		projectList = append(projectList, types.ProjectInfo{
			ProjectId:    item.ProjectId,
			ProjectCode:  item.ProjectCode,
			Title:        item.Title,
			Description:  item.Description,
			Category:     item.Category,
			Status:       item.Status,
			Progress:     item.Progress,
			CreateTime:   item.CreateTime,
			UpdateTime:   item.UpdateTime,
			LastActivity: item.LastActivity,
			Tags:         item.Tags,
		})
	}

	resp = &types.GetProjectListResp{
		List:     projectList,
		Total:    rpcResp.Total,
		PageSize: rpcResp.PageSize,
		Page:     rpcResp.Page,
	}

	return resp, nil
}
