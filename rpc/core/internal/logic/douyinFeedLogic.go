package logic

import (
	"context"
	oss "douyin/rpc/core/internal/OSSClient"
	"douyin/rpc/core/internal/svc"
	"douyin/rpc/core/pb"
	"douyin/rpc/core/utils"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinFeedLogic {
	return &DouyinFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinFeedLogic) DouyinFeed(in *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {
	// todo: add your logic here and delete this line
	_ = InitVideoList(l.svcCtx, l.ctx)

	return &pb.DouyinFeedResponse{
		StatusCode: 0,
		VideoList:  oss.VideoList,
		NextTime:   time.Now().Unix(),
	}, nil
}

func InitVideoList(svcCtx *svc.ServiceContext, ctx context.Context) error {
	// 从数据库中加载videolist
	result, err := svcCtx.VideoModel.FindAll(ctx)
	if err != nil {
		fmt.Println("查询所有video出错了")
		fmt.Println(err)
		return err
	}

	for _, video := range result {
		user, err := svcCtx.UserModel.FindOne(ctx, video.Author)
		if err != nil {
			fmt.Println("FindOneByToken 出错了")
			fmt.Println(err)
			return err
		}
		author, _ := utils.UserModelPb(user)
		pbVideo, _ := utils.VideoModelPb(video, author)
		oss.VideoList = append(oss.VideoList, pbVideo)
	}
	return nil
}
