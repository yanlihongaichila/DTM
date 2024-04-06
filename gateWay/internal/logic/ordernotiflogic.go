package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"gateWay/internal/consts"
	"gateWay/internal/pkg"
	"net/url"
	"order/order"

	"gateWay/internal/svc"
	"gateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderNotifLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderNotifLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderNotifLogic {
	return &OrderNotifLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderNotifLogic) OrderNotif(req url.Values) (resp *types.OrderNotifResponse, err error) {
	verify, err := pkg.VerifySign(l.svcCtx, req)
	if err != nil {
		return nil, err
	}
	orderNo := verify.OutTradeNo
	updateOrder, err := l.svcCtx.OrderCon.UpdateOrder(l.ctx, &order.UpdateOrderRequest{Info: &order.OrderInfo{
		OrderNO: orderNo,
		State:   consts.PAYMENT_SUCCESS,
	}})
	if err != nil {
		return nil, err
	}

	fmt.Println("******************")
	fmt.Println(updateOrder)
	marshal, err := json.Marshal(updateOrder)
	if err != nil {
		return nil, err
	}
	return &types.OrderNotifResponse{Message: string(marshal)}, nil
}
