package douyinMessage

import (
	"context"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageActionLogic {
	return &DouyinMessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinMessageActionLogic) DouyinMessageAction(req *types.DouyinMessageActionRequest) (resp *types.DouyinMessageActionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
