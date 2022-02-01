package main

import (
	"context"
	"github.com/llzzrrr1997/bamboo/framework/gin"
	"github.com/llzzrrr1997/bamboo/framework/middleware"
	"github.com/llzzrrr1997/bamboo/provider/demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())
	_ = core.Bind(&demo.DemoServiceProvider{})
	//core.Use(middleware.Timeout(3 * time.Second))
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8888",
	}
	// 这个goroutine是启动服务的goroutine
	go func() {
		_ = server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
