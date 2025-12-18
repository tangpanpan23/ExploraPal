package project

import (
	"net/http"

	"explorapal/app/api/internal/logic/project"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取项目详情
func GetProjectDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProjectDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewGetProjectDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetProjectDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
