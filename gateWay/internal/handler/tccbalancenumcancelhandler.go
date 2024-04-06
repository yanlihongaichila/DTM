package handler

import (
	"net/http"

	"gateWay/internal/logic"
	"gateWay/internal/svc"
	"gateWay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TCCBalanceNumCancelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TCCBalanceNumCancelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTCCBalanceNumCancelLogic(r.Context(), svcCtx)
		resp, err := l.TCCBalanceNumCancel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
