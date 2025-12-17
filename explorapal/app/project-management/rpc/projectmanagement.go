package main

import (
	"flag"
	"fmt"

	"explorapal/app/project-management/rpc/internal/config"
	"explorapal/app/project-management/rpc/internal/server"
	"explorapal/app/project-management/rpc/internal/svc"
	"explorapal/app/project-management/rpc/projectmanagement"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/projectmanagement.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		projectmanagement.RegisterProjectManagementServiceServer(grpcServer, server.NewProjectManagementServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
