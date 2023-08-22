// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.2
// source: social.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Social_DouyinRelationAction_FullMethodName       = "/douyin.social.social/douyin_relation_action"
	Social_DouyinRelationFollowList_FullMethodName   = "/douyin.social.social/douyin_relation_follow_list"
	Social_DouyinRelationFollowerList_FullMethodName = "/douyin.social.social/douyin_relation_follower_list"
	Social_DouyinRelationFriendList_FullMethodName   = "/douyin.social.social/douyin_relation_friend_list"
	Social_DouyinMessageChat_FullMethodName          = "/douyin.social.social/douyin_message_chat"
	Social_DouyinMessageAction_FullMethodName        = "/douyin.social.social/douyin_message_action"
)

// SocialClient is the client API for Social service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialClient interface {
	DouyinRelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error)
	DouyinRelationFollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error)
	DouyinRelationFollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error)
	DouyinRelationFriendList(ctx context.Context, in *DouyinRelationFriendListRequest, opts ...grpc.CallOption) (*DouyinRelationFriendListResponse, error)
	DouyinMessageChat(ctx context.Context, in *DouyinMessageChatRequest, opts ...grpc.CallOption) (*DouyinMessageChatResponse, error)
	DouyinMessageAction(ctx context.Context, in *DouyinMessageActionRequest, opts ...grpc.CallOption) (*DouyinMessageActionResponse, error)
}

type socialClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialClient(cc grpc.ClientConnInterface) SocialClient {
	return &socialClient{cc}
}

func (c *socialClient) DouyinRelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error) {
	out := new(DouyinRelationActionResponse)
	err := c.cc.Invoke(ctx, Social_DouyinRelationAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) DouyinRelationFollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error) {
	out := new(DouyinRelationFollowListResponse)
	err := c.cc.Invoke(ctx, Social_DouyinRelationFollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) DouyinRelationFollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error) {
	out := new(DouyinRelationFollowerListResponse)
	err := c.cc.Invoke(ctx, Social_DouyinRelationFollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) DouyinRelationFriendList(ctx context.Context, in *DouyinRelationFriendListRequest, opts ...grpc.CallOption) (*DouyinRelationFriendListResponse, error) {
	out := new(DouyinRelationFriendListResponse)
	err := c.cc.Invoke(ctx, Social_DouyinRelationFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) DouyinMessageChat(ctx context.Context, in *DouyinMessageChatRequest, opts ...grpc.CallOption) (*DouyinMessageChatResponse, error) {
	out := new(DouyinMessageChatResponse)
	err := c.cc.Invoke(ctx, Social_DouyinMessageChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) DouyinMessageAction(ctx context.Context, in *DouyinMessageActionRequest, opts ...grpc.CallOption) (*DouyinMessageActionResponse, error) {
	out := new(DouyinMessageActionResponse)
	err := c.cc.Invoke(ctx, Social_DouyinMessageAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialServer is the server API for Social service.
// All implementations must embed UnimplementedSocialServer
// for forward compatibility
type SocialServer interface {
	DouyinRelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error)
	DouyinRelationFollowList(context.Context, *DouyinRelationFollowListRequest) (*DouyinRelationFollowListResponse, error)
	DouyinRelationFollowerList(context.Context, *DouyinRelationFollowerListRequest) (*DouyinRelationFollowerListResponse, error)
	DouyinRelationFriendList(context.Context, *DouyinRelationFriendListRequest) (*DouyinRelationFriendListResponse, error)
	DouyinMessageChat(context.Context, *DouyinMessageChatRequest) (*DouyinMessageChatResponse, error)
	DouyinMessageAction(context.Context, *DouyinMessageActionRequest) (*DouyinMessageActionResponse, error)
	mustEmbedUnimplementedSocialServer()
}

// UnimplementedSocialServer must be embedded to have forward compatible implementations.
type UnimplementedSocialServer struct {
}

func (UnimplementedSocialServer) DouyinRelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DouyinRelationAction not implemented")
}
func (UnimplementedSocialServer) DouyinRelationFollowList(context.Context, *DouyinRelationFollowListRequest) (*DouyinRelationFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DouyinRelationFollowList not implemented")
}
func (UnimplementedSocialServer) DouyinRelationFollowerList(context.Context, *DouyinRelationFollowerListRequest) (*DouyinRelationFollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DouyinRelationFollowerList not implemented")
}
func (UnimplementedSocialServer) DouyinRelationFriendList(context.Context, *DouyinRelationFriendListRequest) (*DouyinRelationFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DouyinRelationFriendList not implemented")
}
func (UnimplementedSocialServer) DouyinMessageChat(context.Context, *DouyinMessageChatRequest) (*DouyinMessageChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DouyinMessageChat not implemented")
}
func (UnimplementedSocialServer) DouyinMessageAction(context.Context, *DouyinMessageActionRequest) (*DouyinMessageActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DouyinMessageAction not implemented")
}
func (UnimplementedSocialServer) mustEmbedUnimplementedSocialServer() {}

// UnsafeSocialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialServer will
// result in compilation errors.
type UnsafeSocialServer interface {
	mustEmbedUnimplementedSocialServer()
}

func RegisterSocialServer(s grpc.ServiceRegistrar, srv SocialServer) {
	s.RegisterService(&Social_ServiceDesc, srv)
}

func _Social_DouyinRelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).DouyinRelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_DouyinRelationAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).DouyinRelationAction(ctx, req.(*DouyinRelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_DouyinRelationFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).DouyinRelationFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_DouyinRelationFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).DouyinRelationFollowList(ctx, req.(*DouyinRelationFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_DouyinRelationFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).DouyinRelationFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_DouyinRelationFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).DouyinRelationFollowerList(ctx, req.(*DouyinRelationFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_DouyinRelationFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).DouyinRelationFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_DouyinRelationFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).DouyinRelationFriendList(ctx, req.(*DouyinRelationFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_DouyinMessageChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinMessageChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).DouyinMessageChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_DouyinMessageChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).DouyinMessageChat(ctx, req.(*DouyinMessageChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_DouyinMessageAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinMessageActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).DouyinMessageAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_DouyinMessageAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).DouyinMessageAction(ctx, req.(*DouyinMessageActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Social_ServiceDesc is the grpc.ServiceDesc for Social service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Social_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.social.social",
	HandlerType: (*SocialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "douyin_relation_action",
			Handler:    _Social_DouyinRelationAction_Handler,
		},
		{
			MethodName: "douyin_relation_follow_list",
			Handler:    _Social_DouyinRelationFollowList_Handler,
		},
		{
			MethodName: "douyin_relation_follower_list",
			Handler:    _Social_DouyinRelationFollowerList_Handler,
		},
		{
			MethodName: "douyin_relation_friend_list",
			Handler:    _Social_DouyinRelationFriendList_Handler,
		},
		{
			MethodName: "douyin_message_chat",
			Handler:    _Social_DouyinMessageChat_Handler,
		},
		{
			MethodName: "douyin_message_action",
			Handler:    _Social_DouyinMessageAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "social.proto",
}
