package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"gozero_micro_food/usermanage/model"
	"gozero_micro_food/usermanage/rpc/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
