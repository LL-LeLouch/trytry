package homestayOrder

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/order/cmd/rpc/order"
	"trytry/common/ctxdata"
	"trytry/common/tool"
	"trytry/common/xerr"

	"trytry/app/order/cmd/api/internal/svc"
	"trytry/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderListLogic) UserHomestayOrderList(req *types.UserHomestayOrderListReq) (*types.UserHomestayOrderListResp, error) {

	userid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.OrderRpc.UserHomestayOrderList(l.ctx, &order.UserHomestayOrderListReq{
		LastId:      req.LastId,
		PageSize:    req.PageSize,
		UserId:      userid,
		TraderState: req.TradeState,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get user homestay order list"), "Failed to get user homestay order list err : %v ,req:%+v", err, req)
	}

	var tyHomestayOrderList []types.UserHomestayOrderListView
	if len(resp.List) > 0 {
		for _, homestayOrder := range resp.List {
			var tyHomestayOrder types.UserHomestayOrderListView
			_ = copier.Copy(&tyHomestayOrder, homestayOrder)
			tyHomestayOrder.OrderTotalPrice = tool.Fen2Yuan(homestayOrder.OrderTotalPrice)
			tyHomestayOrderList = append(tyHomestayOrderList, tyHomestayOrder)
		}
	}
	return &types.UserHomestayOrderListResp{
		List: tyHomestayOrderList,
	}, nil
}
