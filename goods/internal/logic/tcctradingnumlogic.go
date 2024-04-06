package logic

import (
	"context"
	"goods/pkg"

	"goods/goods"
	"goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTCCTradingNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingNumLogic {
	return &TCCTradingNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TCC
func (l *TCCTradingNumLogic) TCCTradingNum(in *goods.TCCTradingNumRequest) (*goods.TCCTradingNumResponse, error) {

	req := map[int64]int64{}

	for _, val := range in.TradingNum {
		req[val.ID] = -val.Num
	}

	err := pkg.TCCTradingNum(req)
	if err != nil {
		return nil, err
	}
	return &goods.TCCTradingNumResponse{}, nil
}
