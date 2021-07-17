package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"gozero_micro_food/usermanage/api/internal/config"
	"gozero_micro_food/usermanage/rpc/user/user"
)

type ServiceContext struct {
	Config config.Config
	User   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   user.NewUser(zrpc.MustNewClient(c.User)),
	}
}
