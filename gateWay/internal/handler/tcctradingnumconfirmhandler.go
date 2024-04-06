package handler

import (
	"context"
	"goods/goods"
	"net/http"

	"gateWay/internal/logic"
	"gateWay/internal/svc"
	"gateWay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 确认冻结
func TCCTradingNumConfirmHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//接收参数
		r.ParseForm()

		var req types.TCCTradingNumConfirm
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//userID := req.UserID // 获取 JSON 中的 userID 字段值
		tccInfos := req.TCCInfos
		reqInfos := []*goods.TCCInfo{}
		for _, val := range tccInfos {
			info := goods.TCCInfo{
				ID:  val.GoodID,
				Num: val.Num,
			}
			reqInfos = append(reqInfos, &info)
		}
		//调用rpc
		_, err := svcCtx.GoodsCon.TCCTradingLockNum(context.Background(), &goods.TCCTradingLockNumRequest{TradingNum: reqInfos})
		if err != nil {
			return
		}

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTCCTradingNumConfirmLogic(r.Context(), svcCtx)
		resp, err := l.TCCTradingNumConfirm(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
