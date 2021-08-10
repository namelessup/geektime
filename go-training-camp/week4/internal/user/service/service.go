package service

import (
	"context"

	v1 "go-geektime/api/user/v1"
	"go-geektime/internal/user/biz"

	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var UserSet = wire.NewSet(NewUserService)

type UserService struct {
	ub *biz.UserBiz
}

func NewUserService(ub *biz.UserBiz) (*UserService, func(), error) {
	return &UserService{ub: ub}, nil, nil
}

func (u *UserService) CreateUser(ctx context.Context, request *v1.UserRequest) (*v1.UserReply, error) {
	if request.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is empty")
	}

	_, err := u.ub.CreateUser(ctx, ToBizUserDO(request))
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return new(v1.UserReply), nil
}

func ToBizUserDO(u1 *v1.UserRequest) *biz.UserDO {
	if u1 == nil {
		return nil
	}

	return &biz.UserDO{
		Username:    u1.Username,
		Password:    u1.Password,
		CreatedTime: u1.CreatedTime.AsTime(),
	}
}
