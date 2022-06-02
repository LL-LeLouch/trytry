package homestayComment

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/travel/model"
	"trytry/common/xerr"

	"trytry/app/travel/cmd/api/internal/svc"
	"trytry/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (*types.CommentListResp, error) {

	whereBuilder := l.svcCtx.HomestayCommentModel.RowBuilder().Where(squirrel.Eq{
		"homestay_id": req.HomestayId,
	})
	commentsList, err := l.svcCtx.HomestayCommentModel.FindAll(l.ctx, whereBuilder, "star_desc")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get activity homestay id set fail rowType: %s ,err : %v", model.HomestayActivityPreferredType, err)
	}
	var resp []types.HomestayComment
	if len(commentsList) > 0 {
		for _, item := range commentsList {
			var tyHomestayComment types.HomestayComment
			_ = copier.Copy(&tyHomestayComment, item)
			resp = append(resp, tyHomestayComment)
		}
	}

	return &types.CommentListResp{
		List: resp,
	}, nil
}
