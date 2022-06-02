package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"trytry/app/travel/cmd/rpc/internal/config"
	"trytry/app/travel/model"
)

type ServiceContext struct {
	Config        config.Config
	HomestayModel model.HomestayModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		HomestayModel: model.NewHomestayModel(conn, c.Cache),
	}
}
