// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package article

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/cms/cmd/api/internal/logic/article"
	"looklook/app/cms/cmd/api/internal/types"
	"net/http"

	"looklook/app/cms/cmd/api/internal/svc"
)

// 获取文章列表
func ArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//调试阶段：直接返回成功文本
		/*httpx.OkJsonCtx(r.Context(), w, map[string]string{
			"message": "请求成功",
		})*/

		var req types.ArticleListReq

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewArticleListLogic(r.Context(), svcCtx)
		resp, err := l.ArticleList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
