package logic

import (
	"context"
	"order/order"

	"gateWay/internal/svc"
	"gateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCBalanceNumConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTCCBalanceNumConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCBalanceNumConfirmLogic {
	return &TCCBalanceNumConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TCCBalanceNumConfirmLogic) TCCBalanceNumConfirm(req *types.TCCBalanceNumConfirmRequest) (resp *types.TCCBalanceNumConfirmResponse, err error) {
	orderNo := req.OrderNo
	tradingOrder, err := l.svcCtx.OrderCon.TCCTradingOrder(l.ctx, &order.TCCTradingOrderRequest{Info: &order.OrderInfo{
		OrderNO: orderNo,
	}})
	if err != nil {
		return nil, err
	}
	return &types.TCCBalanceNumConfirmResponse{OrderNo: tradingOrder.Info.OrderNO}, nil
}
