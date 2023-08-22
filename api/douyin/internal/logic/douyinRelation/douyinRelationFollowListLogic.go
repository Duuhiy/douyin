package douyinRelation

import (
	"context"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFollowListLogic {
	return &DouyinRelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationFollowListLogic) DouyinRelationFollowList(req *types.DouyinRelationFollowListRequest) (resp *types.DouyinRelationFollowListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
