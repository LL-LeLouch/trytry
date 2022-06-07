package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"trytry/app/order/cmd/rpc/internal/config"
	"trytry/app/order/model"
	"trytry/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config      config.Config
	AsynqClient *asynq.Client

	TravelRpc travel.Travel

	HomestayOrderModel model.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		AsynqClient: newAsynqClient(c),

		TravelRpc: travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),

		HomestayOrderModel: model.NewHomestayOrderModel(conn, c.Cache),
	}
}
