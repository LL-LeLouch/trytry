package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/order/model"
	"trytry/common/xerr"

	"trytry/app/order/cmd/rpc/internal/svc"
	"trytry/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户民宿订单
func (l *UserHomestayOrderListLogic) UserHomestayOrderList(in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	whereBuilder := l.svcCtx.HomestayOrderModel.RowBuilder().Where(squirrel.Eq{"user_id": in.UserId})
	//过滤器中有支持的状态，否则全部返回
	if in.TraderState >= model.HomestayOrderTradeStateCancel && in.TraderState <= model.HomestayOrderTradeStateExpire {
		whereBuilder = whereBuilder.Where(squirrel.Eq{"trade_state": in.TraderState})
	}
	list, err := l.svcCtx.HomestayOrderModel.FindPageListByIdDESC(l.ctx, whereBuilder, in.LastId, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR failed to get user's homestay order err : %v , in :%+v", err, in)
	}

	var resp []*pb.HomestayOrder
	if len(list) > 0 {
		for _, HomestayOrder := range list {
			var pbHomestayOrder pb.HomestayOrder
			_ = copier.Copy(&pbHomestayOrder, HomestayOrder)
			resp = append(resp, &pbHomestayOrder)
		}
	}
	return &pb.UserHomestayOrderListResp{
		List: resp,
	}, nil
}
