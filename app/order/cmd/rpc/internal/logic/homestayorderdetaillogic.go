package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/common/xerr"

	"trytry/app/order/cmd/rpc/internal/svc"
	"trytry/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayOrderDetailLogic {
	return &HomestayOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomestayOrderDetail 民宿订单详情
func (l *HomestayOrderDetailLogic) HomestayOrderDetail(in *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {
	homeOrder, err := l.svcCtx.HomestayOrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayOrderModel  FindOneBySn pb err : %v , sn : %s", err, in.Sn)
	}
	var resp pb.HomestayOrder
	if homeOrder != nil {
		copier.Copy(&resp, homeOrder)

		resp.CreateTime = homeOrder.CreateTime.Unix()
		resp.LiveStartDate = homeOrder.LiveStartDate.Unix()
		resp.LiveEndDate = homeOrder.LiveEndDate.Unix()

	}

	return &pb.HomestayOrderDetailResp{
		HomestayOrder: &resp,
	}, nil
}
