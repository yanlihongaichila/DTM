package logic

import (
	"context"
	"order/order"

	"gateWay/internal/svc"
	"gateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCBalanceNumTryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTCCBalanceNumTryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCBalanceNumTryLogic {
	return &TCCBalanceNumTryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TCCBalanceNumTryLogic) TCCBalanceNumTry(req *types.TCCBalanceNumTryRequest) (resp *types.TCCBalanceNumTryResponse, err error) {
	userId := req.UserID
	tradingOrder, err := l.svcCtx.OrderCon.TCCTradingOrder(l.ctx, &order.TCCTradingOrderRequest{Info: &order.OrderInfo{
		UserID: userId,
	}})
	if err != nil {
		return nil, err
	}
	return &types.TCCBalanceNumTryResponse{ID: tradingOrder.Info.ID}, nil
}
