package logic

import (
	"context"
	"fmt"
	"gateWay/internal/consts"
	"gateWay/internal/pkg"
	"gateWay/internal/svc"
	"gateWay/internal/types"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"goods/goods"
	"goods/goodsclient"
	"math/rand"
	"order/order"
	"time"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	//先查询商品,判断商品是否存在,并判断库存是否够
	//传入商品的map
	reqGoodsMap := make(map[int64]int64)
	for _, val := range req.Goods {
		reqGoodsMap[val.GoodID] = val.Num
	}
	ids := []int64{}
	for _, val := range req.Goods {
		ids = append(ids, val.GoodID)
	}
	reqInfos, err := l.svcCtx.GoodsCon.GetGoodsByIDs(l.ctx, &goods.GetGoodsByIDsRequest{IDs: ids})
	if err != nil {
		return nil, err
	}
	//得到仓库中库存的map
	shopStockMap := make(map[int64]int64)
	for _, val := range reqInfos.Infos {
		shopStockMap[val.ID] = val.Stock
	}
	fmt.Println("****************")
	fmt.Println(shopStockMap)
	for key, val := range shopStockMap {
		if val-reqGoodsMap[key] < 0 {
			return nil, fmt.Errorf("%v商品库存不足", key)
		}
		reqGoodsMap[key] = -reqGoodsMap[key]
	}
	//修改库存
	//_, err = l.updatedStock(reqGoodsMap)
	//if err != nil {
	//	return nil, err
	//}

	updatedReq := []*goods.UpdateStockReq{}

	for key, val := range reqGoodsMap {
		conInfo := goods.UpdateStockReq{
			ID:  key,
			Num: val,
		}

		updatedReq = append(updatedReq, &conInfo)
	}
	//获取商品总价
	amount, err := l.getGoodAmount(req)
	if err != nil {
		return nil, err
	}
	orderInfo := order.OrderInfo{
		UserID:    req.UserID,
		Amount:    amount.String(),
		State:     consts.PAYMENT_WAIT,
		CreatedAt: time.Now().Unix(),
		OrderNO:   fmt.Sprintf("%v%v%v%v%v", time.Now().UnixNano(), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)),
	}
	//createdOrder, err := l.svcCtx.OrderCon.CreateOrder(l.ctx, &order.CreateOrderRequest{Info: &orderInfo})
	//if err != nil {
	//	return nil, err
	//}
	//sage
	gid := uuid.NewString()

	//sage := dtmgrpc.NewSagaGrpc("127.0.0.1:36790", gid).Add("10.2.171.14:7750/goods.Goods/UpdateGoodsStocks", "10.2.171.14:7750/goods.Goods/UpdateGoodsStocksRollback", &goods.UpdateGoodsStocksRequest{GoodsInfos: updatedReq}).Add("10.2.171.14:7751/order.Order/CreateOrder", "10.2.171.14:7751/order.Order/GetOrderRollback", &order.CreateOrderRequest{Info: &orderInfo})

	trandinsInfos := []*goods.TCCInfo{}
	for key, val := range reqGoodsMap {
		aaa := goods.TCCInfo{
			ID:  key,
			Num: val,
		}
		trandinsInfos = append(trandinsInfos, &aaa)
	}

	//tcc
	err = dtmgrpc.TccGlobalTransaction("127.0.0.1:36790", gid, func(tcc *dtmgrpc.TccGrpc) error {
		err = tcc.CallBranch(&goods.TCCTradingNumRequest{TradingNum: trandinsInfos}, "10.2.171.14:7750/goods.Goods/TCCTradingNum", "10.2.171.14:7750/goods.Goods/TCCTradingLockNum", "10.2.171.14:7750/goods.Goods/TCCTradingRollbackNum", &goods.TCCTradingLockNumResponse{})
		if err != nil {
			fmt.Println("22222")
			fmt.Println(err)
			return err
		}

		err = tcc.CallBranch(&order.TCCTradingOrderRequest{Info: &orderInfo}, "10.2.171.14:7751/order.Order/TCCTradingOrder", "10.2.171.14:7751/order.Order/TCCTradingLockOrder", "10.2.171.14:7751/order.Order/TCCTradingRollbackOrder", &order.TCCTradingRollbackOrderResponse{})
		if err != nil {
			fmt.Println("1111111111111111111111111")
			fmt.Println(err)
			return err
		}
		return nil
	})
	//添加订单商品表
	orderGoodsInfos := []*order.OrderGoodsInfo{}

	//sage提交
	//err = sage.Submit()
	if err != nil {
		return nil, err
	}
	for _, val := range reqInfos.Infos {
		info := order.OrderGoodsInfo{
			//todo: 有瑕疵
			OrderID:   int64(rand.Intn(1000)),
			GoodsID:   val.ID,
			UnitPrice: val.Price,
			GoodName:  val.Name,
			Num:       -reqGoodsMap[val.ID],
		}

		orderGoodsInfos = append(orderGoodsInfos, &info)
	}
	_, err = l.svcCtx.OrderCon.CreateOrderGoods(l.ctx, &order.CreateOrderGoodsRequest{Infos: orderGoodsInfos})
	if err != nil {
		return nil, err
	}

	url, err := pkg.GetWebPayUrl(l.svcCtx, "商品", orderInfo.OrderNO, amount.String())
	if err != nil {
		return nil, err
	}

	return &types.CreateOrderResponse{Url: url}, nil
}

// 得总价
func (l *OrderCreateLogic) getGoodAmount(req *types.CreateOrderRequest) (decimal.Decimal, error) {
	var goodIDs []int64
	for _, val := range req.Goods {
		goodIDs = append(goodIDs, val.GoodID)
	}

	reqInfos, err := l.svcCtx.GoodsCon.GetGoodsByIDs(l.ctx, &goods.GetGoodsByIDsRequest{IDs: goodIDs})
	if err != nil {
		return decimal.Decimal{}, err
	}

	amount := decimal.NewFromInt(0)

	goodMap := make(map[int64]*goods.GoodsInfo)
	for _, val := range reqInfos.Infos {
		goodMap[val.ID] = val
	}

	for _, val := range req.Goods {
		goodInfo, ok := goodMap[val.GoodID]
		if !ok {
			return decimal.Decimal{}, fmt.Errorf("商品不存在")
		}

		unitPrice, err := decimal.NewFromString(goodInfo.Price)
		if err != nil {
			return decimal.Decimal{}, err
		}
		amount = amount.Add(unitPrice.Mul(decimal.NewFromInt(val.Num)))
	}

	return amount, nil
}

// 修改库存
func (l *OrderCreateLogic) updatedStock(upInfos map[int64]int64) (*goodsclient.UpdateGoodsStocksResponse, error) {
	updateGoodsInfos := goods.UpdateGoodsStocksRequest{}
	updatedReq := updateGoodsInfos.GoodsInfos

	for key, val := range upInfos {
		conInfo := goods.UpdateStockReq{
			ID:  key,
			Num: val,
		}

		updatedReq = append(updatedReq, &conInfo)
	}

	stocks, err := l.svcCtx.GoodsCon.UpdateGoodsStocks(l.ctx, &goods.UpdateGoodsStocksRequest{GoodsInfos: updatedReq})
	if err != nil {
		return nil, err
	}

	return stocks, nil
}
