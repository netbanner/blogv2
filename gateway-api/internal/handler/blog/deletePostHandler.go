// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"net/http"

	"blogV2/gateway-api/internal/logic/blog"
	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除文章
func DeletePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostById
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewDeletePostLogic(r.Context(), svcCtx)
		err := l.DeletePost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
