package main

import (
	v1 "go-geektime/api/user/v1"
	"go-geektime/internal/user/service"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	v1.RegisterUserServiceServer(grpcServer, &service.UserServer{})
}

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
