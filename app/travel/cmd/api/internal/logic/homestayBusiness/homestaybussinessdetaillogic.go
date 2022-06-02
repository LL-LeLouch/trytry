package homestayBusiness

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/usercenter/cmd/rpc/usercenter"
	"trytry/app/usercenter/model"
	"trytry/common/xerr"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessDetailLogic {
	return &HomestayBusinessDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessDetailLogic) HomestayBusinessDetail(req *types.HomestayBusinessDetailReq) (*types.HomestayBusinessDetailResp, error) {
	homestayBusiness, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayBusinessDetail findOne  id %d  failed: %v", req.Id, err)
	}
	var tyHomestayBusiness types.HomestayBusinessBoss
	if homestayBusiness != nil {
		userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
			Uid: req.Id,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("get boss info fail"), "get boss info fail ,  userId : %d ,err:%v", homestayBusiness.UserId, err)
		}
		if userResp.User != nil && userResp.User.Id > 0 {
			_ = copier.Copy(&tyHomestayBusiness, userResp.User)
		}
	}
	return &types.HomestayBusinessDetailResp{
		Boss: tyHomestayBusiness,
	}, nil
}
