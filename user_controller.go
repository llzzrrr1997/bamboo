package main

import (
	"github.com/llzzrrr1997/bamboo/framework/gin"
	"github.com/llzzrrr1997/bamboo/provider/demo"
)

func UserLoginController(c *gin.Context) {
	// 等待10s才结束执行
	s := c.MustMake(demo.Key).(demo.Service)
	data := s.GetFoo()
	_ = c.ISetOkStatus().IJson(data)
}
