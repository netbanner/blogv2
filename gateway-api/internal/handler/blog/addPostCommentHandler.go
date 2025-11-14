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

// 添加评论
func AddPostCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPostComment
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewAddPostCommentLogic(r.Context(), svcCtx)
		err := l.AddPostComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
