package douyinRelation

import (
	"context"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFollowerListLogic {
	return &DouyinRelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationFollowerListLogic) DouyinRelationFollowerList(req *types.DouyinRelationFollowerListRequest) (resp *types.DouyinRelationFollowerListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
