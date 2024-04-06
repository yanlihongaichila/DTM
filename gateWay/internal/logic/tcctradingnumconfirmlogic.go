package logic

import (
	"context"

	"gateWay/internal/svc"
	"gateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingNumConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTCCTradingNumConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingNumConfirmLogic {
	return &TCCTradingNumConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TCCTradingNumConfirmLogic) TCCTradingNumConfirm(req *types.TCCTradingNumConfirm) (resp *types.TCCTradingNumConfirm, err error) {
	// todo: add your logic here and delete this line

	return
}
