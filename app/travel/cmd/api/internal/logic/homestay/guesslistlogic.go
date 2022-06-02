package homestay

import (
	"context"
	"github.com/jinzhu/copier"
	"trytry/common/tool"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GuessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuessListLogic {
	return &GuessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuessListLogic) GuessList(req *types.GuessListReq) (*types.GuessListResp, error) {
	var resp []types.Homestay

	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, l.svcCtx.HomestayModel.RowBuilder(), 0, 5)
	if err != nil {
		return nil, err
	}
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

	return &types.GuessListResp{List: resp}, nil
}
