package douyinFavorite

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/douyinFavorite"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinFavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinFavoriteActionRequest
		vars := r.URL.Query()
		if token, ok := vars["token"]; ok {
			req.Token = token[0]
		} else {
			fmt.Println("token错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("token错误"))
		}

		if videoIdStr, ok := vars["video_id"]; ok {
			videoId, err := strconv.ParseInt(videoIdStr[0], 10, 64)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, errors.New("video_id格式错误"))
			}
			req.Video_id = videoId
		} else {
			fmt.Println("video_id错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id错误"))
		}

		if ActionTypeStr, ok := vars["action_type"]; ok {
			ActionType, err := strconv.ParseInt(ActionTypeStr[0], 10, 64)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, errors.New("action_type格式错误"))
			}
			req.Action_type = int32(ActionType)
		} else {
			fmt.Println("action_type错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("action_type错误"))
		}

		l := douyinFavorite.NewDouyinFavoriteActionLogic(r.Context(), svcCtx)
		fmt.Println("调用 api DouyinFavoriteAction")
		resp, err := l.DouyinFavoriteAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
