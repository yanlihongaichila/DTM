package logic

import (
	"context"
	"goods/pkg"

	"goods/goods"
	"goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingLockNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTCCTradingLockNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingLockNumLogic {
	return &TCCTradingLockNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TCCTradingLockNumLogic) TCCTradingLockNum(in *goods.TCCTradingLockNumRequest) (*goods.TCCTradingLockNumResponse, error) {
	req := map[int64]int64{}

	for _, val := range in.TradingNum {
		req[val.ID] = -val.Num
	}

	err := pkg.TCCTradingLockNum(req)
	if err != nil {
		return nil, err
	}

	return &goods.TCCTradingLockNumResponse{}, nil
}
