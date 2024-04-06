package svc

import (
	"gateWay/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"goods/goodsclient"
	"order/orderclient"
)

type ServiceContext struct {
	Config   config.Config
	GoodsCon goodsclient.Goods
	OrderCon orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GoodsCon: goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsClient)),
		OrderCon: orderclient.NewOrder(zrpc.MustNewClient(c.OrderClient)),
	}
}
