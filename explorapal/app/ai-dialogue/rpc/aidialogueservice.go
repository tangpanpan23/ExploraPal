package main

import (
	"flag"
	"fmt"

	"explorapal/app/ai-dialogue/rpc/internal/config"
	"explorapal/app/ai-dialogue/rpc/internal/server"
	"explorapal/app/ai-dialogue/rpc/internal/svc"
	"explorapal/app/ai-dialogue/rpc/aidialogue"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ai-dialogue.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		aidialogue.RegisterAIDialogueServiceServer(grpcServer, server.NewAIDialogueServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("🚀 Starting AI dialogue rpc server at %s...\n", c.ListenOn)
	s.Start()
}
