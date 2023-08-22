// Code generated by goctl. DO NOT EDIT.
// Source: social.proto

package server

import (
	"context"

	"douyin/rpc/social/internal/logic"
	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"
)

type CoreServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCoreServer
}

func NewCoreServer(svcCtx *svc.ServiceContext) *CoreServer {
	return &CoreServer{
		svcCtx: svcCtx,
	}
}

func (s *CoreServer) DouyinRelationAction(ctx context.Context, in *pb.DouyinRelationActionRequest) (*pb.DouyinRelationActionResponse, error) {
	l := logic.NewDouyinRelationActionLogic(ctx, s.svcCtx)
	return l.DouyinRelationAction(in)
}

func (s *CoreServer) DouyinRelationFollowList(ctx context.Context, in *pb.DouyinRelationFollowListRequest) (*pb.DouyinRelationFollowListResponse, error) {
	l := logic.NewDouyinRelationFollowListLogic(ctx, s.svcCtx)
	return l.DouyinRelationFollowList(in)
}

func (s *CoreServer) DouyinRelationFollowerList(ctx context.Context, in *pb.DouyinRelationFollowerListRequest) (*pb.DouyinRelationFollowerListResponse, error) {
	l := logic.NewDouyinRelationFollowerListLogic(ctx, s.svcCtx)
	return l.DouyinRelationFollowerList(in)
}

func (s *CoreServer) DouyinRelationFriendList(ctx context.Context, in *pb.DouyinRelationFriendListRequest) (*pb.DouyinRelationFriendListResponse, error) {
	l := logic.NewDouyinRelationFriendListLogic(ctx, s.svcCtx)
	return l.DouyinRelationFriendList(in)
}

func (s *CoreServer) DouyinMessageChat(ctx context.Context, in *pb.DouyinMessageChatRequest) (*pb.DouyinMessageChatResponse, error) {
	l := logic.NewDouyinMessageChatLogic(ctx, s.svcCtx)
	return l.DouyinMessageChat(in)
}

func (s *CoreServer) DouyinMessageAction(ctx context.Context, in *pb.DouyinMessageActionRequest) (*pb.DouyinMessageActionResponse, error) {
	l := logic.NewDouyinMessageActionLogic(ctx, s.svcCtx)
	return l.DouyinMessageAction(in)
}
