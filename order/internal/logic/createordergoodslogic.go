package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderGoodsLogic {
	return &CreateOrderGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderGoodsLogic) CreateOrderGoods(in *order.CreateOrderGoodsRequest) (*order.CreateOrderGoodsResponse, error) {
	// todo: add your logic here and delete this line
	err := pkg.CreateOrderGoods(in.Infos)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderGoodsResponse{}, nil
}
