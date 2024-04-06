package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTCCTradingOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingOrderLogic {
	return &TCCTradingOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TCC
func (l *TCCTradingOrderLogic) TCCTradingOrder(in *order.TCCTradingOrderRequest) (*order.TCCTradingOrderResponse, error) {
	TCCTradingOrderInfo, err := pkg.TCCTradingOrder(in.Info)
	if err != nil {
		return nil, err
	}

	resInfo := order.OrderInfo{
		ID:        int64(TCCTradingOrderInfo.ID),
		UserID:    TCCTradingOrderInfo.UserID,
		OrderNO:   TCCTradingOrderInfo.OrderNO,
		Amount:    TCCTradingOrderInfo.Amount,
		State:     TCCTradingOrderInfo.State,
		CreatedAt: TCCTradingOrderInfo.CreatedAt.Unix(),
	}

	return &order.TCCTradingOrderResponse{Info: &resInfo}, nil
}
