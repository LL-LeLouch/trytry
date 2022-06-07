package homestayOrder

import (
	"context"
	"github.com/pkg/errors"
	"trytry/app/order/cmd/api/internal/svc"
	"trytry/app/order/cmd/api/internal/types"
	"trytry/app/order/cmd/rpc/order"
	"trytry/app/travel/cmd/rpc/pb"
	"trytry/app/travel/model"
	"trytry/common/ctxdata"
	"trytry/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHomestayOrderLogic) CreateHomestayOrder(req *types.CreateHomestayOrderReq) (*types.CreateHomestayOrderResp, error) {
	homestayResp, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &pb.HomestayDetailReq{
		Id: req.HomestayId,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR  rpc TravelRpc.HomestayDetail  errors: %v , req :%+v", err, req)
	}
	if (homestayResp.Homestay == nil && homestayResp.Homestay.Id == 0) || err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("homestay no exists"), "CreateHomestayOrder homestay no exists id : %d", req.HomestayId)
	}
	userId := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.OrderRpc.CreateHomestayOrder(l.ctx, &order.CreateHomestayOrderReq{
		HomestayId:    req.HomestayId,
		IsFood:        req.IsFood,
		LiveStartTime: req.LiveStartTime,
		LiveEndTime:   req.LiveEndTime,
		UserId:        userId,
		LivePeopleNum: req.LivePeopleNum,
		Remark:        req.Remark,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("create homestay order failed"), "ERROR create homestay order rpc CreateHomestayOrder fail req: %+v , err : %v ", req, err)
	}

	return &types.CreateHomestayOrderResp{
		OrderSn: resp.Sn,
	}, nil
}
