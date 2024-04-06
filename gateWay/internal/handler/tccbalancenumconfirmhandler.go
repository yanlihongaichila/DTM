package handler

import (
	"net/http"

	"gateWay/internal/logic"
	"gateWay/internal/svc"
	"gateWay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TCCBalanceNumConfirmHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TCCBalanceNumConfirmRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTCCBalanceNumConfirmLogic(r.Context(), svcCtx)
		resp, err := l.TCCBalanceNumConfirm(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
