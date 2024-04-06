package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingRollbackOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTCCTradingRollbackOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingRollbackOrderLogic {
	return &TCCTradingRollbackOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TCCTradingRollbackOrderLogic) TCCTradingRollbackOrder(in *order.TCCTradingRollbackOrderRequest) (*order.TCCTradingRollbackOrderResponse, error) {
	_, err := pkg.TCCTradingRollbackOrder(in.Info.OrderNO)
	if err != nil {
		return nil, err
	}
	return &order.TCCTradingRollbackOrderResponse{}, nil
}
