package expression

import (
	"net/http"

	"explorapal/app/api/internal/logic/expression"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 语音转文字
func SpeechToTextHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SpeechToTextReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := expression.NewSpeechToTextLogic(r.Context(), svcCtx)
		resp, err := l.SpeechToText(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
