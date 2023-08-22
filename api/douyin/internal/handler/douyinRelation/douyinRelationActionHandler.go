package douyinRelation

import (
	"net/http"

	"douyin/api/douyin/internal/logic/douyinRelation"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinRelationActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinRelationActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := douyinRelation.NewDouyinRelationActionLogic(r.Context(), svcCtx)
		resp, err := l.DouyinRelationAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
