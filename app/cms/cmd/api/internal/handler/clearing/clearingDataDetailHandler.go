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

// 获取出清数据详情
func ClearingDataDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClearingDataDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := clearing.NewClearingDataDetailLogic(r.Context(), svcCtx)
		resp, err := l.ClearingDataDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
