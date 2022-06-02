package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"trytry/app/order/cmd/rpc/order"
	"trytry/app/payment/cmd/api/internal/config"
	"trytry/app/payment/cmd/rpc/payment"
	"trytry/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config

	PaymentRpc    payment.Payment
	OrderRpc      order.Order
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		PaymentRpc:    payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
