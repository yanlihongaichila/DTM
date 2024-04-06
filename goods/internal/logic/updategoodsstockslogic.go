package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"goods/goods"
	"goods/internal/svc"
	"goods/pkg"
)

type UpdateGoodsStocksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsStocksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsStocksLogic {
	return &UpdateGoodsStocksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsStocksLogic) UpdateGoodsStocks(in *goods.UpdateGoodsStocksRequest) (*goods.UpdateGoodsStocksResponse, error) {
	req := map[int64]int64{}

	for _, val := range in.GoodsInfos {
		req[val.ID] = val.Num
	}

	err := pkg.UpdateGoodsStocks(req)
	if err != nil {
		return nil, err
	}
	return &goods.UpdateGoodsStocksResponse{}, nil
}
