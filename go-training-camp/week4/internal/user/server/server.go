package server

import (
	v1 "go-geektime/api/user/v1"
	"go-geektime/internal/user/service"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

var UserSet = wire.NewSet(NewUserGPRCServer)

func NewUserGPRCServer(user *service.UserService) (*grpc.Server, func(), error) {
	srv := grpc.NewServer()
	v1.RegisterUserServiceServer(srv, user)
	return srv, func() {
		srv.Stop()
	}, nil
}
