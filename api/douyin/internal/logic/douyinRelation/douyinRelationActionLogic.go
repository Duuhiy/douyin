package douyinRelation

import (
	"context"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationActionLogic {
	return &DouyinRelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationActionLogic) DouyinRelationAction(req *types.DouyinRelationActionRequest) (resp *types.DouyinRelationActionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
