package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	User zrpc.RpcClientConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}