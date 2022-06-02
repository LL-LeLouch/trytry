package homestay

import (
	"context"
	"github.com/jinzhu/copier"
	"trytry/app/travel/cmd/rpc/travel"
	"trytry/common/tool"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {
	homestayResp, err := l.svcCtx.TraceClient.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var tyHomestay types.Homestay
	if homestayResp != nil {
		_ = copier.Copy(&tyHomestay, homestayResp)

		tyHomestay.FoodPrice = tool.Fen2Yuan(homestayResp.Homestay.FoodPrice)
		tyHomestay.HomestayPrice = tool.Fen2Yuan(homestayResp.Homestay.HomestayPrice)
		tyHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestayResp.Homestay.MarketHomestayPrice)
	}

	return
}
