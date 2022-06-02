package homestayBusiness

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"trytry/app/travel/model"
	"trytry/app/usercenter/cmd/rpc/usercenter"
	"trytry/common/xerr"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodBossLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodBossLogic {
	return &GoodBossLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodBossLogic) GoodBoss(req *types.GoodBossReq) (*types.GoodBossResp, error) {
	whereBuilder := l.svcCtx.HomestayActivityModel.RowBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityGoodBusiType,
		"row_status": model.HomestayActivityUpStatus,
	})
	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, 0, 10, "data_id desc")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get GoodBoss pb err. rowType: %s ,err : %v", model.HomestayActivityGoodBusiType, err)
	}
	var resp []types.HomestayBusinessBoss
	if len(homestayActivityList) > 0 {
		err := mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, homestayActivity := range homestayActivityList {
				source <- homestayActivity.DataId
			}
		}, func(item interface{}, writer mr.Writer, cancel func(error)) {
			id := item.(int64)
			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			if err != nil && err != model.ErrNotFound {
				logx.WithContext(l.ctx).Errorf("ActivityHomestayListLogic ActivityHomestayList 获取活动数据失败 id : %d ,err : %v", id, err)
				return
			}

			userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
				Uid: homestay.UserId,
			})
			if err != nil {
				logx.WithContext(l.ctx).Errorf("GoodListLogic GoodList fail userId : %d ,err:%v", homestay.UserId, err)
				return
			}
			if userResp.User != nil && userResp.User.Id > 0 {
				writer.Write(userResp.User)
			}

		}, func(pipe <-chan interface{}, cancel func(error)) {
			for item := range pipe {
				var tyHomestayBusiness types.HomestayBusinessBoss
				_ = copier.Copy(&tyHomestayBusiness, item)

				// compute star todo
				//Rank 排名没有

				resp = append(resp, tyHomestayBusiness)
			}
		})
		if err != nil {
			return nil, err
		}
	}

	return &types.GoodBossResp{
		List: resp,
	}, nil
}
