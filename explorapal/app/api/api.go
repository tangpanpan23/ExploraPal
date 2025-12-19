package main

import (
	"flag"
	"fmt"

	"explorapal/app/api/internal/config"
	"explorapal/app/api/internal/handler"
	"explorapal/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 注册路由处理器
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("🚀 Starting API server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("📋 API文档: http://%s:%d/api/common/ping\n", c.Host, c.Port)
	server.Start()
}
