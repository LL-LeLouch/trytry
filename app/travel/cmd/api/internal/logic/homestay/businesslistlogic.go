package homestay

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/travel/model"
	"trytry/common/tool"
	"trytry/common/xerr"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.BusinessListReq) (*types.BusinessListResp, error) {
	whereBuilder := l.svcCtx.HomestayModel.RowBuilder().Where(squirrel.Eq{
		"homestay_business_id": req.HomestayBusinessId,
	})
	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR  FindPageListByIdDESC HomestayBusinessId: %d ,err %v", req.HomestayBusinessId, err)
	}

	var resp []types.Homestay
	if len(list) > 0 {
		for _, homestay := range list {
			var tyHomestay types.Homestay
			_ = copier.Copy(&tyHomestay, homestay)

			tyHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			tyHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			tyHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)

			resp = append(resp, tyHomestay)

		}
	}

	return &types.BusinessListResp{List: resp}, nil
}
