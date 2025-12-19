package common

import (
	"net/http"

	"explorapal/app/api/internal/logic/common"
	"explorapal/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 健康检查
func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := common.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
