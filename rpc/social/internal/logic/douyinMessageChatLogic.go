package logic

import (
	"context"

	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageChatLogic {
	return &DouyinMessageChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinMessageChatLogic) DouyinMessageChat(in *pb.DouyinMessageChatRequest) (*pb.DouyinMessageChatResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DouyinMessageChatResponse{}, nil
}
