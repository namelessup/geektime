package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	g, ctx := errgroup.WithContext(context.Background())
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "ok")
	})

	srv := &http.Server{
		Addr:         ":8881",
		WriteTimeout: time.Second * 3,
		Handler:      srvMux,
	}
	g.Go(func() error {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("http服务异常退出, err=%+v", err)
			return err
		}

		log.Println("http服务正常退出")
		return err
	})

	g.Go(func() error {
		_ = server1(ctx)
		log.Println("server1进程退出!")
		return nil
	})

	s2 := SleepServer{}
	g.Go(func() error {
		return s2.server2()
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit
		log.Println("接收到正常退出信号！开始关闭服务")
		_ = srv.Shutdown(ctx)
		log.Println("http服务关闭完成")

		s2.close()
		log.Println("server2服务关闭完成")
		return nil
	})

	err := g.Wait()
	if err != nil && errors.Is(err, context.Canceled) {
		log.Printf("程序异常退出!, err=%+v", err)
		os.Exit(1)
	}

	log.Println("主线程正常退出!")
	os.Exit(0)
}

func server1(ctx context.Context) error {
	select {
	case <-ctx.Done():
		time.Sleep(5 * time.Second)
		return ctx.Err()
	}
}

type SleepServer struct {
	interval time.Duration
	goon     atomic.Value
}

func (s *SleepServer) server2() error {
	val := s.goon.Load()
	goon, ok := val.(bool)
	if !ok {
		return errors.New("继续标识设置错误！")
	}

	for goon {
		time.Sleep(time.Second)
	}
	return nil
}

func (s *SleepServer) close() {
	s.goon.Store(false)
}
