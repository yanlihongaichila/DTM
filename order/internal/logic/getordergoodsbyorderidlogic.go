package logic

import (
	"context"
	"order/internal/pkg"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderGoodsByOrderIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderGoodsByOrderIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderGoodsByOrderIDLogic {
	return &GetOrderGoodsByOrderIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderGoodsByOrderIDLogic) GetOrderGoodsByOrderID(in *order.GetOrderGoodsByOrderIDRequest) (*order.GetOrderGoodsByOrderIDResponse, error) {
	// todo: add your logic here and delete this line
	id, err := pkg.GetOrderGoodsByOrderID(in.OrderID)
	if err != nil {
		return nil, err
	}

	res := &order.GetOrderGoodsByOrderIDResponse{}
	resInfos := res.Infos

	for _, val := range id {
		info := order.OrderGoodsInfo{
			ID:        int64(val.ID),
			OrderID:   val.OrderID,
			GoodsID:   val.GoodsID,
			UnitPrice: val.UnitPrice,
			GoodName:  val.GoodName,
			Num:       val.Num,
		}

		resInfos = append(resInfos, &info)
	}
	return &order.GetOrderGoodsByOrderIDResponse{Infos: resInfos}, nil
}
