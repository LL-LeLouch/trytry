package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"trytry/app/travel/cmd/api/internal/config"
	"trytry/app/travel/cmd/rpc/travel"
	"trytry/app/travel/model"
	"trytry/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	TraceClient   travel.Travel
	UsercenterRpc usercenter.Usercenter

	//model
	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel
	HomestayCommentModel  model.HomestayCommentModel

	//RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                c,
		TraceClient:           travel.NewTravel(zrpc.MustNewClient(c.TravelRpc)),
		UsercenterRpc:         usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpc)),
		HomestayModel:         model.NewHomestayModel(conn, c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(conn, c.Cache),
		HomestayBusinessModel: model.NewHomestayBusinessModel(conn, c.Cache),
		HomestayCommentModel:  model.NewHomestayCommentModel(conn, c.Cache),
	}
}
