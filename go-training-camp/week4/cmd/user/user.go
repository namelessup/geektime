package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go-geektime/internal/user/conf"

	_ "github.com/mattn/go-sqlite3"
)

var configPath = flag.String("config_path", "../configs/user.yaml", "setting file path")

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

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		clean()
	}()

	listener, err := net.Listen("tcp", conf.Config.Grpc.Addr)
	if err != nil {
		panic(err)
	}

	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}
