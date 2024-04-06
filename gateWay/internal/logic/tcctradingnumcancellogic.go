package logic

import (
	"context"

	"gateWay/internal/svc"
	"gateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TCCTradingNumCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTCCTradingNumCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TCCTradingNumCancelLogic {
	return &TCCTradingNumCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TCCTradingNumCancelLogic) TCCTradingNumCancel(req *types.TCCTradingNumCancel) (resp *types.TCCTradingNumCancel, err error) {
	// todo: add your logic here and delete this line

	return
}
