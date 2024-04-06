package handler

import (
	"fmt"
	"net/http"

	"gateWay/internal/logic"
	"gateWay/internal/svc"
	"gateWay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderNotifHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		var req types.OrderNotifRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderNotifLogic(r.Context(), svcCtx)
		resp, err := l.OrderNotif(r.Form)
		fmt.Println("OrderNotif Err ******************************")
		fmt.Println(err)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
