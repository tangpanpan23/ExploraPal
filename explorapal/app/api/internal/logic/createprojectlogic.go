package logic

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
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

func (l *CreateProjectLogic) CreateProject(req *types.CreateProjectReq) (resp *types.CreateProjectResp, err error) {
	// 调用项目管理RPC服务
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:9001"}, // 项目管理RPC服务地址
	})
	if err != nil {
		l.Logger.Errorf("创建RPC客户端失败: %v", err)
		return nil, err
	}

	projectClient := projectmanagement.NewProjectManagementServiceClient(client.Conn())

	rpcResp, err := projectClient.CreateProject(l.ctx, &projectmanagement.CreateProjectReq{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Tags:        req.Tags,
	})
	if err != nil {
		l.Logger.Errorf("调用项目管理RPC服务失败: %v", err)
		return nil, err
	}

	resp = &types.CreateProjectResp{
		ProjectId:   rpcResp.ProjectId,
		ProjectCode: rpcResp.ProjectCode,
		Status:      "active", // 默认状态
	}

	return resp, nil
}
