package http

import (
	"github.com/llzzrrr1997/bamboo/app/http/module/demo"
	"github.com/llzzrrr1997/bamboo/framework/gin"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
