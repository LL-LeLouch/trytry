package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"trytry/app/mqueue/cmd/job/jobtype"
	"trytry/app/order/model"
	"trytry/common/xerr"

	"trytry/app/order/cmd/rpc/internal/svc"
	"trytry/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayOrderTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayOrderTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayOrderTradeStateLogic {
	return &UpdateHomestayOrderTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomestayOrderTradeState 更新民宿订单状态
func (l *UpdateHomestayOrderTradeStateLogic) UpdateHomestayOrderTradeState(in *pb.UpdateHomestayOrderTradeStateReq) (*pb.UpdateHomestayOrderTradeStateResp, error) {
	// 1、查看当前订单
	homestayOrder, err := l.svcCtx.HomestayOrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateHomestayOrderTradeState FindOneBySn  failed: %v in:%+v ", err, in)
	}
	if homestayOrder == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("HomestayOrder not exist"), "order no exists  in : %+v", in)
	}
	if homestayOrder.TradeState == in.TradeState {
		return &pb.UpdateHomestayOrderTradeStateResp{}, nil
	}

	// 2、验证订单状态
	if err := l.verifyOrderTradeState(in.TradeState, homestayOrder.TradeState); err != nil {
		return nil, errors.WithMessagef(err, " , in : %+v", in)
	}

	//3、通知user
	if in.TradeState == model.HomestayOrderTradeStateWaitUse {
		payload, err := json.Marshal(jobtype.PaySuccessNotifyUserPayload{
			Order: homestayOrder,
		})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("pay success notify user task json Marshal fail, err :%+v , sn : %s", err, homestayOrder.Sn)
		} else {
			_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgPaySuccessNotifyUser, payload))
			if err != nil {
				logx.WithContext(l.ctx).Errorf("pay success notify user  insert queue fail err :%+v , sn : %s", err, homestayOrder.Sn)
			}
		}
	}
	return &pb.UpdateHomestayOrderTradeStateResp{
		Id:              homestayOrder.Id,
		UserId:          homestayOrder.UserId,
		Sn:              homestayOrder.Sn,
		TradeCode:       homestayOrder.TradeCode,
		Title:           homestayOrder.Title,
		LiveStartDate:   homestayOrder.LiveStartDate.Unix(),
		LiveEndDate:     homestayOrder.LiveEndDate.Unix(),
		OrderTotalPrice: homestayOrder.OrderTotalPrice,
	}, nil
}

//验证
func (l *UpdateHomestayOrderTradeStateLogic) verifyOrderTradeState(newTradeState, oldTradeState int64) error {
	if newTradeState == model.HomestayOrderTradeStateWaitPay {
		return errors.Wrapf(xerr.NewErrMsg("Changing this status is not supported"),
			"Changing this status is not supported newTradeState: %d, oldTradeState: %d",
			newTradeState,
			oldTradeState)
	}
	if newTradeState == model.HomestayOrderTradeStateCancel {

		if oldTradeState != model.HomestayOrderTradeStateWaitPay {
			return errors.Wrapf(xerr.NewErrMsg("只有待支付的订单才能被取消"),
				"Only orders pending payment can be cancelled newTradeState: %d, oldTradeState: %d",
				newTradeState,
				oldTradeState)
		} else if newTradeState == model.HomestayOrderTradeStateWaitUse {
			if oldTradeState != model.HomestayOrderTradeStateWaitPay {
				return errors.Wrapf(xerr.NewErrMsg("Only orders pending payment can change this status"),
					"Only orders pending payment can change this status newTradeState: %d, oldTradeState: %d",
					newTradeState,
					oldTradeState)
			}
		} else if newTradeState == model.HomestayOrderTradeStateRefund {
			if oldTradeState != model.HomestayOrderTradeStateWaitUse {
				return errors.Wrapf(xerr.NewErrMsg("Only unused orders can be changed to this status"),
					"Only unused orders can be changed to this status newTradeState: %d, oldTradeState: %d",
					newTradeState,
					oldTradeState)
			}
		} else if newTradeState == model.HomestayOrderTradeStateExpire {
			if oldTradeState != model.HomestayOrderTradeStateWaitUse {
				return errors.Wrapf(xerr.NewErrMsg("Only unused orders can be changed to this status"),
					"Only unused orders can be changed to this status newTradeState: %d, oldTradeState: %d",
					newTradeState,
					oldTradeState)
			}
		}

	}
	return nil
}
