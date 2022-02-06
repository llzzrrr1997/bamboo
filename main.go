package main

import (
	"github.com/llzzrrr1997/bamboo/app/console"
	"github.com/llzzrrr1997/bamboo/app/http"
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/provider/app"
	"github.com/llzzrrr1997/bamboo/framework/provider/kernel"
)

func main() {
	/*// 创建engine结构
	core := gin.New()
	// 绑定具体的服务
	core.Bind(&app.BambooAppProvider{})
	core.Bind(&demo.DemoProvider{})

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	bambooHttp.Routes(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
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
	}*/
	// 初始化服务容器
	container := framework.NewBambooContainer()
	// 绑定App服务提供者
	container.Bind(&app.BambooAppProvider{})
	// 后续初始化需要绑定的服务提供者...

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.BambooKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
