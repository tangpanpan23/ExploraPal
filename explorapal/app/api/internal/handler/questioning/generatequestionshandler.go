package questioning

import (
	"net/http"

	"explorapal/app/api/internal/logic/questioning"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 生成引导问题
func GenerateQuestionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateQuestionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := questioning.NewGenerateQuestionsLogic(r.Context(), svcCtx)
		resp, err := l.GenerateQuestions(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
