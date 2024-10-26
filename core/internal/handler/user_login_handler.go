package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tree-cloud/core/internal/logic"
	"tree-cloud/core/internal/svc"
	"tree-cloud/core/internal/types"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
