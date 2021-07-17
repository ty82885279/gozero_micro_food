package handler

import (
	"gozero_micro_food/lee"
	"gozero_micro_food/usermanage/api/internal/logic"
	"net/http"

	"gozero_micro_food/usermanage/api/internal/svc"
	"gozero_micro_food/usermanage/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserinfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, lee.FailureResponse(nil, err.Error(), 1000))
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(req)
		if err != nil {
			httpx.OkJson(w, lee.FailureResponse(nil, err.Error(), 1000))
		} else {
			httpx.OkJson(w, lee.SuccessResponse(resp, "请求成功"))
		}
	}
}
