package project

import (
	"net/http"

	"explorapal/app/api/internal/logic/project"
	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新项目状态
func UpdateProjectStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProjectStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewUpdateProjectStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProjectStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
