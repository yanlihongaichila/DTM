package main

import (
	"flag"
	"fmt"
	"goods/goods"
	"goods/internal/config"
	"goods/internal/server"
	"goods/internal/svc"
	"goods/pkg"
	"pkg/app"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goods.yaml", "the config file")

func main() {
	flag.Parse()
	//连接mysql
	err := app.Init("mysql")
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建表
	err = pkg.Migrate()
	if err != nil {
		fmt.Println(err)
		return
	}
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods.RegisterGoodsServer(grpcServer, server.NewGoodsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
