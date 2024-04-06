package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	createOrder, err := pkg.CreateOrder(in.Info)
	if err != nil {
		return nil, err
	}
	resInfo := order.OrderInfo{
		ID:        int64(createOrder.ID),
		UserID:    createOrder.UserID,
		OrderNO:   createOrder.OrderNO,
		Amount:    createOrder.Amount,
		State:     createOrder.State,
		CreatedAt: createOrder.CreatedAt.Unix(),
	}

	return &order.CreateOrderResponse{Info: &resInfo}, nil
}
