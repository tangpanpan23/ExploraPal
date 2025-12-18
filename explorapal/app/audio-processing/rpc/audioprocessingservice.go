package main

import (
	"flag"
	"fmt"

	"explorapal/app/audio-processing/rpc/internal/config"
	"explorapal/app/audio-processing/rpc/internal/server"
	"explorapal/app/audio-processing/rpc/internal/svc"
	"explorapal/app/audio-processing/rpc/audioprocessing"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/audio-processing.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		audioprocessing.RegisterAudioProcessingServiceServer(grpcServer, server.NewAudioProcessingServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("🎵 Starting audio processing rpc server at %s...\n", c.ListenOn)
	s.Start()
}
