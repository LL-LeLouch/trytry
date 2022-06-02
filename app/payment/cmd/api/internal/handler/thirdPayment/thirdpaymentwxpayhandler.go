package thirdPayment

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"trytry/common/result"

	"trytry/app/payment/cmd/api/internal/logic/thirdPayment"
	"trytry/app/payment/cmd/api/internal/svc"
	"trytry/app/payment/cmd/api/internal/types"
)

func ThirdPaymentWxPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThirdPaymentWxPayReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := thirdPayment.NewThirdPaymentWxPayLogic(r.Context(), svcCtx)
		resp, err := l.ThirdPaymentWxPay(&req)
		result.HttpResult(r, w, resp, err)
	}
}
