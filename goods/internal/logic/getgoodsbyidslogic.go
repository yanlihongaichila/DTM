package logic

import (
	"context"
	"goods/goods"
	"goods/internal/svc"
	"goods/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsByIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIDsLogic {
	return &GetGoodsByIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsByIDsLogic) GetGoodsByIDs(in *goods.GetGoodsByIDsRequest) (*goods.GetGoodsByIDsResponse, error) {
	// todo: add your logic here and delete this line
	ds, err := pkg.GetGoodsByIDs(in.IDs)
	if err != nil {
		return nil, err
	}
	goodsRes := goods.GetGoodsByIDsResponse{}
	goodsInfos := goodsRes.Infos

	for _, val := range ds {
		info := goods.GoodsInfo{
			ID:    int64(val.ID),
			Name:  val.Name,
			Price: val.Price,
			Stock: val.Stock,
		}
		goodsInfos = append(goodsInfos, &info)
	}
	return &goods.GetGoodsByIDsResponse{Infos: goodsInfos}, nil
}
