package main

import (
	"flag"
	"fmt"

	"gozero_micro_food/usermanage/rpc/user/internal/config"
	"gozero_micro_food/usermanage/rpc/user/internal/server"
	"gozero_micro_food/usermanage/rpc/user/internal/svc"
	"gozero_micro_food/usermanage/rpc/user/pro/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
