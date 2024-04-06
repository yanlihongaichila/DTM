package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *order.UpdateOrderRequest) (*order.UpdateOrderResponse, error) {
	// todo: add your logic here and delete this line
	err := pkg.UpdateOrder(in.Info)
	if err != nil {
		return nil, err
	}
	return &order.UpdateOrderResponse{}, nil
}
