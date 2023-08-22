package douyinRelation

import (
	"context"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFriendListLogic {
	return &DouyinRelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationFriendListLogic) DouyinRelationFriendList(req *types.DouyinRelationFriendListRequest) (resp *types.DouyinRelationFriendListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
