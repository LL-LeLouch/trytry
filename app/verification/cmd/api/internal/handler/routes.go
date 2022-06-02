// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"trytry/app/verification/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/verification/email",
				Handler: verifyemailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/verification/image",
				Handler: verifyimageHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}