package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/travel/model"
	"trytry/common/xerr"

	"trytry/app/travel/cmd/rpc/internal/svc"
	"trytry/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// homestayDetail
func (l *HomestayDetailLogic) HomestayDetail(in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayDetail pb failed id %d", in.Id)
	}
	var pbHomestay pb.Homestay
	if homestay != nil {
		_ = copier.Copy(&pbHomestay, homestay)
	}

	return &pb.HomestayDetailResp{
		Homestay: &pbHomestay,
	}, nil
}
