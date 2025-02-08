package main

import (
	"flag"
	"fmt"

	"github.com/solate/admin_backend/app/admin/internal/config"
	"github.com/solate/admin_backend/app/admin/internal/handler"
	"github.com/solate/admin_backend/app/admin/internal/middleware"
	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/pkg/common/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	// 加载配置文件
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 配置log
	c.LoadLogConf()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.LoggerMiddleware)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 使用拦截器，处理返回值
	httpx.SetOkHandler(response.OkHanandler)
	httpx.SetErrorHandlerCtx(response.ErrHandler(c.Name))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
