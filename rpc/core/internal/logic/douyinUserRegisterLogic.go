package logic

import (
	"context"
	"douyin/rpc/core/pb"
	"fmt"

	"douyin/model/user"
	"douyin/rpc/core/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinUserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinUserRegisterLogic {
	return &DouyinUserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinUserRegisterLogic) DouyinUserRegister(in *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入 DouyinUserRegister rpc 服务")
	user := &user.User{
		Name:     in.Username,
		Password: in.Password,
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return &pb.DouyinUserRegisterResponse{
			StatusCode: 1,
		}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &pb.DouyinUserRegisterResponse{
			StatusCode: 1,
		}, err
	}
	return &pb.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     id,
	}, nil
}
