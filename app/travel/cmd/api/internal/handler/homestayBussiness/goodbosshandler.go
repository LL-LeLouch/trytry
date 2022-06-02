package homestayBussiness

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"trytry/common/result"

	"trytry/app/travel/cmd/api/internal/logic/homestayBusiness"
	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"
)

func GoodBossHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodBossReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBusiness.NewGoodBossLogic(r.Context(), svcCtx)
		resp, err := l.GoodBoss(&req)
		result.HttpResult(r, w, resp, err)
	}
}
