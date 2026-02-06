// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/cms/cmd/api/internal/logic/clearing"
	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"
)

// 创建出清数据
func CreateClearingDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateClearingDataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := clearing.NewCreateClearingDataLogic(r.Context(), svcCtx)
		resp, err := l.CreateClearingData(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
