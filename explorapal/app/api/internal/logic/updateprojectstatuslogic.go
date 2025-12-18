package logic

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type UpdateProjectStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectStatusLogic {
	return &UpdateProjectStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProjectStatusLogic) UpdateProjectStatus(req *types.UpdateProjectStatusReq) (resp *types.CommonStatusResp, err error) {
	// 调用项目管理RPC服务
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9001"}, // 项目管理RPC服务地址
	})
	if err != nil {
		l.Logger.Errorf("创建RPC客户端失败: %v", err)
		return nil, err
	}

	projectClient := projectmanagement.NewProjectManagementServiceClient(client.Conn())

	_, err = projectClient.UpdateProjectStatus(l.ctx, &projectmanagement.UpdateProjectStatusReq{
		ProjectId: req.ProjectId,
		UserId:    req.UserId,
		Status:    req.Status,
	})
	if err != nil {
		l.Logger.Errorf("调用项目管理RPC服务失败: %v", err)
		return nil, err
	}

	resp = &types.CommonStatusResp{
		Code:    200,
		Message: "更新成功",
	}

	return resp, nil
}
