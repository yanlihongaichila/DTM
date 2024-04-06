package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderRollbackLogic {
	return &GetOrderRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderRollbackLogic) GetOrderRollback(in *order.GetOrderRollbackRequest) (*order.GetOrderRollbackResponse, error) {

	_, err := pkg.CreateOrderRollback(in.Info.ID)
	if err != nil {
		return nil, err
	}
	//return &order.GetOrderResponse{Info: &resInfo}, nil
	return &order.GetOrderRollbackResponse{}, nil
}
