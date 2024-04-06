package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingLockOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTCCTradingLockOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingLockOrderLogic {
	return &TCCTradingLockOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TCCTradingLockOrderLogic) TCCTradingLockOrder(in *order.TCCTradingLockOrderRequest) (*order.TCCTradingLockOrderResponse, error) {
	tCCTradingLockOrder, err := pkg.TCCTradingLockOrder(in.Info)
	if err != nil {
		return nil, err
	}
	resInfo := order.OrderInfo{
		ID:        int64(tCCTradingLockOrder.ID),
		UserID:    tCCTradingLockOrder.UserID,
		OrderNO:   tCCTradingLockOrder.OrderNO,
		Amount:    tCCTradingLockOrder.Amount,
		State:     tCCTradingLockOrder.State,
		CreatedAt: tCCTradingLockOrder.CreatedAt.Unix(),
	}

	return &order.TCCTradingLockOrderResponse{Info: &resInfo}, nil
}
