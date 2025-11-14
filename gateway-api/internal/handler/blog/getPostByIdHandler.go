// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"net/http"

	"blogV2/gateway-api/internal/logic/blog"
	"blogV2/gateway-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 根据文章id查询文章
func GetPostByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := blog.NewGetPostByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetPostById()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
