package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tree-cloud/core/internal/logic"
	"tree-cloud/core/internal/svc"
	"tree-cloud/core/internal/types"
)

func MailCodeSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMailCodeSendLogic(r.Context(), svcCtx)
		resp, err := l.MailCodeSend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
