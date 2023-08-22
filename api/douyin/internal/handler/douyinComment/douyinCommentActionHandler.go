package douyinComment

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/douyinComment"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinCommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinCommentActionRequest
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

		if actionTypeStr, ok := vars["action_type"]; ok {
			actionType, err := strconv.ParseInt(actionTypeStr[0], 10, 32)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, errors.New("video_id格式错误"))
			}
			req.Action_type = int32(actionType)
		} else {
			fmt.Println("video_id错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id错误"))
		}

		if commentText, ok := vars["comment_text"]; ok {
			req.Comment_text = commentText[0]
		} else {
			fmt.Println("video_id错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id错误"))
		}

		if commentIdStr, ok := vars["comment_id"]; ok {
			commentId, err := strconv.ParseInt(commentIdStr[0], 10, 64)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, errors.New("video_id格式错误"))
			}
			req.Video_id = commentId
		} else {
			fmt.Println("video_id错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id错误"))
		}

		l := douyinComment.NewDouyinCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.DouyinCommentAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
