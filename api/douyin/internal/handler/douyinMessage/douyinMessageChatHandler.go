package douyinMessage

import (
	"net/http"

	"douyin/api/douyin/internal/logic/douyinMessage"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinMessageChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinMessageChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := douyinMessage.NewDouyinMessageChatLogic(r.Context(), svcCtx)
		resp, err := l.DouyinMessageChat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
