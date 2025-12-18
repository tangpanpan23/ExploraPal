package main

import (
	"flag"
	"fmt"

	"explorapal/app/video-processing/rpc/internal/config"
	"explorapal/app/video-processing/rpc/internal/server"
	"explorapal/app/video-processing/rpc/internal/svc"
	"explorapal/app/video-processing/rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/video-processing.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterVideoProcessingServiceServer(grpcServer, server.NewVideoProcessingServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("🎥 Starting video processing rpc server at %s...\n", c.ListenOn)
	s.Start()
}
