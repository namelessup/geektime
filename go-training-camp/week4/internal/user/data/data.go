package data

import (
	"context"

	"go-geektime/internal/user/biz"
	"go-geektime/internal/user/conf"
	"go-geektime/internal/user/data/ent"

	"entgo.io/ent/dialect"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(NewUserData, NewUserRepo)

type UserData struct {
	client *ent.Client
}

// TODO 也可以依赖配置
func NewUserData() (data *UserData, clean func(), err error) {
	client, err := ent.Open(dialect.SQLite, conf.Config.SQLite.Addr)
	if err != nil {
		return
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		return
	}

	data = &UserData{client: client}
	clean = func() {
		_ = client.Close()
	}
	return
}

type userRepo struct {
	data *UserData
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.UserDO) (int64, error) {
	u1 := u.data.client.User.Create()
	u1.SetUsername(user.Username)
	u1.SetCreatedTime(user.CreatedTime)
	u1.SetNillablePassword(&user.Password)
	du, err := u1.Save(ctx)
	if err != nil {
		return 0, err
	}

	return du.ID, nil
}

func NewUserRepo(data *UserData) biz.UserRepo {
	return &userRepo{data: data}
}
