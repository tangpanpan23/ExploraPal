package handler

import (
	"net/http"

	"explorapal/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("Health check request from %s", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":200,"msg":"pong","service":"explorapal-api"}`))
	}
}
