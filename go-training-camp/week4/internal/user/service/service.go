package service

import (
	"context"

	v1 "go-geektime/api/user/v1"
)

type UserServer struct {
}

func (u *UserServer) CreateUser(ctx context.Context, request *v1.UserRequest) (*v1.UserReply, error) {
	return nil, nil
}
