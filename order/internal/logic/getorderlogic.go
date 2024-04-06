package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	getOrder, err := pkg.GetOrder(in.ID)
	if err != nil {
		return nil, err
	}

	resInfo := order.OrderInfo{
		ID:        int64(getOrder.ID),
		UserID:    getOrder.UserID,
		OrderNO:   getOrder.OrderNO,
		Amount:    getOrder.Amount,
		State:     getOrder.State,
		CreatedAt: getOrder.CreatedAt.Unix(),
	}
	return &order.GetOrderResponse{Info: &resInfo}, nil
}
