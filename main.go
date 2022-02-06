package main

import (
	"github.com/llzzrrr1997/bamboo/app/console"
	"github.com/llzzrrr1997/bamboo/app/http"
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/provider/app"
	"github.com/llzzrrr1997/bamboo/framework/provider/config"
	"github.com/llzzrrr1997/bamboo/framework/provider/distributed"
	"github.com/llzzrrr1997/bamboo/framework/provider/env"
	"github.com/llzzrrr1997/bamboo/framework/provider/id"
	"github.com/llzzrrr1997/bamboo/framework/provider/kernel"
	"github.com/llzzrrr1997/bamboo/framework/provider/log"
	"github.com/llzzrrr1997/bamboo/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewBambooContainer()
	// 绑定App服务提供者
	container.Bind(&app.BambooAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.BambooEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.BambooConfigProvider{})
	container.Bind(&id.BambooIDProvider{})
	container.Bind(&trace.BambooTraceProvider{})
	container.Bind(&log.BambooLogServiceProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.BambooKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
