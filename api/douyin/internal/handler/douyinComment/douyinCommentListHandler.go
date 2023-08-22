package douyinComment

import (
	"net/http"

	"douyin/api/douyin/internal/logic/douyinComment"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinCommentListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := douyinComment.NewDouyinCommentListLogic(r.Context(), svcCtx)
		resp, err := l.DouyinCommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
