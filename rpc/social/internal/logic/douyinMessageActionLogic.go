package logic

import (
	"context"

	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageActionLogic {
	return &DouyinMessageActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinMessageActionLogic) DouyinMessageAction(in *pb.DouyinMessageActionRequest) (*pb.DouyinMessageActionResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DouyinMessageActionResponse{}, nil
}
