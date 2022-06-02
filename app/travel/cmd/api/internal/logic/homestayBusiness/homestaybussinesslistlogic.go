package homestayBusiness

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/common/xerr"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessListLogic {
	return &HomestayBusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessListLogic) HomestayBusinessList(req *types.HomestayBusinessListReq) (*types.HomestayBusinessListResp, error) {
	whereBuilder := l.svcCtx.HomestayBusinessModel.RowBuilder()
	list, err := l.svcCtx.HomestayBusinessModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayBusinessList FindPageListByIdDESC pb failed:  req : %+v , err:%v", req, err)
	}
	var resp []types.HomestayBusinessListInfo

	//HomestayBusinessListInfo 中没有
	//  SellMonth   //月销售
	//	PersonConsume  //个人消费
	if len(list) > 0 {
		for _, item := range list {
			var tyHomestayBusinessListInfo types.HomestayBusinessListInfo
			_ = copier.Copy(&tyHomestayBusinessListInfo, item)
			resp = append(resp, tyHomestayBusinessListInfo)
		}
	}

	return &types.HomestayBusinessListResp{
		List: resp,
	}, nil
}
