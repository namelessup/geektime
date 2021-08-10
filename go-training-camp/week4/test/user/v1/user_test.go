package v1

import (
	"context"
	"log"
	"testing"
	"time"

	v1 "go-geektime/api/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateUser(t *testing.T) {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := v1.NewUserServiceClient(conn)

	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &v1.UserRequest{
		Username:    "Ben",
		Password:    "123789",
		CreatedTime: timestamppb.Now(),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
