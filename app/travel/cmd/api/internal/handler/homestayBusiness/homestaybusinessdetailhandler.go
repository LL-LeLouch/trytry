package homestayBusiness

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"trytry/common/result"

	"trytry/app/travel/cmd/api/internal/logic/homestayBusiness"
	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"
)

func HomestayBusinessDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBusinessDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBusiness.NewHomestayBusinessDetailLogic(r.Context(), svcCtx)
		resp, err := l.HomestayBusinessDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
