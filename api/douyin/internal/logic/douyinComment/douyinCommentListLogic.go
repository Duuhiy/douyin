package douyinComment

import (
	"context"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinCommentListLogic {
	return &DouyinCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinCommentListLogic) DouyinCommentList(req *types.DouyinCommentListRequest) (resp *types.DouyinCommentListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
