package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"trytry/app/order/cmd/mq/internal/config"
	"trytry/app/order/cmd/rpc/order"
	"trytry/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc      order.Order
	UserCenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
