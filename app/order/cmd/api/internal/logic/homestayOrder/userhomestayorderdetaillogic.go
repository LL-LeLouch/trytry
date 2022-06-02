package homestayOrder

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/order/cmd/rpc/order"
	"trytry/app/order/model"
	"trytry/app/payment/cmd/rpc/payment"
	"trytry/common/ctxdata"
	"trytry/common/tool"
	"trytry/common/xerr"

	"trytry/app/order/cmd/api/internal/svc"
	"trytry/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderDetailLogic {
	return &UserHomestayOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderDetailLogic) UserHomestayOrderDetail(req *types.UserHomestayOrderDetailReq) (*types.UserHomestayOrderDetailResp, error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: req.Sn,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay order detail fail"), " rpc get HomestayOrderDetail err:%v , sn : %s", err, req.Sn)
	}
	var tyOrderDetail types.UserHomestayOrderDetailResp
	if userId == resp.HomestayOrder.UserId {
		copier.Copy(&tyOrderDetail, resp)
		//重置价格.
		tyOrderDetail.OrderTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.OrderTotalPrice)
		tyOrderDetail.FoodTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.FoodTotalPrice)
		tyOrderDetail.HomestayTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.HomestayTotalPrice)
		tyOrderDetail.HomestayPrice = tool.Fen2Yuan(resp.HomestayOrder.HomestayPrice)
		tyOrderDetail.FoodPrice = tool.Fen2Yuan(resp.HomestayOrder.FoodPrice)
		tyOrderDetail.MarketHomestayPrice = tool.Fen2Yuan(resp.HomestayOrder.MarketHomestayPrice)

		//支付信息.
		if resp.HomestayOrder.TradeState != model.HomestayOrderTradeStateCancel && resp.HomestayOrder.TradeState != model.HomestayOrderTradeStateWaitPay {
			paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentSuccessRefundByOrderSn(l.ctx, &payment.GetPaymentSuccessRefundByOrderSnReq{
				OrderSn: resp.HomestayOrder.Sn,
			})
			if err != nil {
				logx.WithContext(l.ctx).Errorf("Failed to get order payment information err : %v , orderSn:%s", err, resp.HomestayOrder.Sn)
			}
			if paymentResp != nil {
				tyOrderDetail.PayTime = paymentResp.PaymentDetail.PayTime
				tyOrderDetail.PayType = paymentResp.PaymentDetail.PayMode
			}
		}
		return &tyOrderDetail, nil
	}
	return nil, nil
}
