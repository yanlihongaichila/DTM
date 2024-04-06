package logic

import (
	"context"

	"goods/goods"
	"goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingRollbackNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTCCTradingRollbackNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingRollbackNumLogic {
	return &TCCTradingRollbackNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TCCTradingRollbackNumLogic) TCCTradingRollbackNum(in *goods.TCCTradingRollbackNumRequest) (*goods.TCCTradingRollbackNumResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.TCCTradingRollbackNumResponse{}, nil
}
