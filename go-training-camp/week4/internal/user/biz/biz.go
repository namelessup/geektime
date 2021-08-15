package biz

import (
	"context"
	"time"

	"github.com/google/wire"
)

var UserSet = wire.NewSet(NewUserBiz)

type UserDO struct {
	Username    string
	Password    string
	CreatedTime time.Time
}

type UserRepo interface {
	CreateUser(ctx context.Context, do *UserDO) (int64, error)
}

type UserBiz struct {
	repo UserRepo
}

func NewUserBiz(repo UserRepo) (*UserBiz, func(), error) {
	return &UserBiz{repo: repo}, func() {}, nil
}

func (ub *UserBiz) CreateUser(ctx context.Context, do *UserDO) (int64, error) {
	return ub.repo.CreateUser(ctx, do)
}
