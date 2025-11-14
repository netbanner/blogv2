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

// 根据用户id查询所以文章
func GetPostByUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewGetPostByUserIdLogic(r.Context(), svcCtx)
		resp, err := l.GetPostByUserId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
