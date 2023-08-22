package douyinMessage

import (
	"context"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageChatLogic {
	return &DouyinMessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinMessageChatLogic) DouyinMessageChat(req *types.DouyinMessageChatRequest) (resp *types.DouyinMessageChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
