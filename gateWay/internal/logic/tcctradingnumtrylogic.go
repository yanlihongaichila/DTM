package logic

import (
	"context"

	"gateWay/internal/svc"
	"gateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingNumTryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTCCTradingNumTryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingNumTryLogic {
	return &TCCTradingNumTryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TCCTradingNumTryLogic) TCCTradingNumTry(req *types.TCCTradingNumTry) (resp *types.TCCTradingNumTry, err error) {
	// todo: add your logic here and delete this line

	return
}
