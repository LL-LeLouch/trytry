package homestay

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"trytry/app/travel/model"
	"trytry/common/tool"
	"trytry/common/xerr"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayListLogic) HomestayList(req *types.HomestayListReq) (*types.HomestayListResp, error) {

	whereBuilder := l.svcCtx.HomestayActivityModel.RowBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityPreferredType,
		"row_status": model.HomestayActivityUpStatus,
	})
	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "data_id desc")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR get activity homestay id set fail rowType: %s ,err : %v", model.HomestayActivityPreferredType, err)
	}

	var resp []types.Homestay
	if len(homestayActivityList) > 0 {
		err := mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, homestayActivity := range homestayActivityList {
				source <- homestayActivity.DataId
			}
		}, func(item interface{}, writer mr.Writer, cancel func(error)) {
			id := item.(int64)

			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)

			if err != nil && err != model.ErrNotFound {
				logx.WithContext(l.ctx).Errorf("ERROR HomestayListLogic ActivityHomestayList 获取活动数据失败 id : %d ,err : %v", id, err)
				return
			}
			writer.Write(homestay)
		}, func(pipe <-chan interface{}, cancel func(error)) {

			for item := range pipe {
				homestay := item.(*model.Homestay)
				var tyHomestay types.Homestay
				_ = copier.Copy(&tyHomestay, homestay)

				tyHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
				tyHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
				tyHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)

				resp = append(resp, tyHomestay)
			}
		})
		if err != nil {
			return nil, err
		}
	}
	return &types.HomestayListResp{List: resp}, nil
}
