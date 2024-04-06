package logic

import (
	"context"
	"goods/pkg"

	"goods/goods"
	"goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsStocksRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsStocksRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsStocksRollbackLogic {
	return &UpdateGoodsStocksRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsStocksRollbackLogic) UpdateGoodsStocksRollback(in *goods.UpdateGoodsStocksRollbackRequest) (*goods.UpdateGoodsStocksRollbackResponse, error) {
	req := map[int64]int64{}

	for _, val := range in.GoodsInfos {
		req[val.ID] = val.Num
	}

	err := pkg.UpdateGoodsStocksRollback(req)
	if err != nil {
		return nil, err
	}
	return &goods.UpdateGoodsStocksRollbackResponse{}, nil
}
