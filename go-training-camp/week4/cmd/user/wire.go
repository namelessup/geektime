// +build wireinject

package main

import (
	"go-geektime/internal/user/biz"
	"go-geektime/internal/user/data"
	"go-geektime/internal/user/service"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func depend() (*grpc.Server, func(), error) {
	wire.Build(service.UserSet, biz.UserSet, data.UserSet, UserSet)
	return &grpc.Server{}, nil, nil
}
