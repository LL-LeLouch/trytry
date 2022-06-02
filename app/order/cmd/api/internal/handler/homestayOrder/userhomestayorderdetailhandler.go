package homestayOrder

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"trytry/common/result"

	"trytry/app/order/cmd/api/internal/logic/homestayOrder"
	"trytry/app/order/cmd/api/internal/svc"
	"trytry/app/order/cmd/api/internal/types"
)

func UserHomestayOrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHomestayOrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayOrder.NewUserHomestayOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserHomestayOrderDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
