package main

import (
	"flag"
	"log"
	"net"

	v1 "go-geektime/api/user/v1"
	"go-geektime/internal/user/conf"
	"go-geektime/internal/user/service"

	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

var configPath = flag.String("config_path", "/Users/kim/CodeHub/geektime/go-training-camp/week4/configs/user.yaml", "setting file path")

func main() {
	if flag.Parsed() {
		flag.Parse()
	}

	err := conf.InitUserConfig(*configPath)
	if err != nil {
		log.Panic(err)
	}

	srv, clean, err := depend()
	if err != nil {
		log.Panic(err)
	}

	defer clean()

	listener, err := net.Listen("tcp", conf.Config.Grpc.Addr)
	if err != nil {
		panic(err)
	}

	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func NewGRPC(user *service.UserService) (*grpc.Server, func(), error) {
	srv := grpc.NewServer()
	v1.RegisterUserServiceServer(srv, user)
	return srv, nil, nil
}

var UserSet = wire.NewSet(NewGRPC)

// func main() {
// // 使用内存中SQLite数据库来创建 ent.Client
// client, err := ent.Open(dialect.SQLite, "file:geektime?mode=memory&cache=shared&_fk=1")
// if err != nil {
// 	log.Fatalf("failed opening connection to sqlite: %v", err)
// }
// defer client.Close()
// ctx := context.Background()
// // 运行自动迁移工具来创建所有Schema资源
// if err := client.Schema.Create(ctx); err != nil {
// 	log.Fatalf("failed creating schema resources: %v", err)
// }
//
// u1 := client.User.Create()
// u1.SetUsername("tom")
// u1.SetCreatedTime(time.Now())
//
// task1, err := u1.Save(ctx)
// if err != nil {
// 	log.Fatalf("failed creating a user: %v", err)
// }
// fmt.Println(task1)
// }
