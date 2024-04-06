package logic

import (
	"context"
	"order/order"

	"gateWay/internal/svc"
	"gateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCBalanceNumCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTCCBalanceNumCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCBalanceNumCancelLogic {
	return &TCCBalanceNumCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TCCBalanceNumCancelLogic) TCCBalanceNumCancel(req *types.TCCBalanceNumCancelRequest) (resp *types.TCCBalanceNumCancelResponse, err error) {
	orderNo := req.OrderNo
	_, err = l.svcCtx.OrderCon.TCCTradingRollbackOrder(l.ctx, &order.TCCTradingRollbackOrderRequest{Info: &order.OrderInfo{
		OrderNO: orderNo,
	}})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
