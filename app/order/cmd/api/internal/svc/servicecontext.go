package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"trytry/app/order/cmd/api/internal/config"
	"trytry/app/order/cmd/rpc/order"
	"trytry/app/payment/cmd/rpc/payment"
	"trytry/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   order.Order
	PaymentRpc payment.Payment
	TravelRpc  travel.Travel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		TravelRpc:  travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
	}
}
