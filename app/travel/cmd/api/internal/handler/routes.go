// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	homestay "trytry/app/travel/cmd/api/internal/handler/homestay"
	homestayBusiness "trytry/app/travel/cmd/api/internal/handler/homestayBusiness"
	homestayComment "trytry/app/travel/cmd/api/internal/handler/homestayComment"
	"trytry/app/travel/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestay/homestayList",
				Handler: homestay.HomestayListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/businessList",
				Handler: homestay.BusinessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/guessList",
				Handler: homestay.GuessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/homestayDetail",
				Handler: homestay.HomestayDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayBusiness/goodBoss",
				Handler: homestayBusiness.GoodBossHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayBusiness/homestayBusinessList",
				Handler: homestayBusiness.HomestayBusinessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayBusiness/homestayBusinessDetail",
				Handler: homestayBusiness.HomestayBusinessDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayComment/commentList",
				Handler: homestayComment.CommentListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)
}
