package main

import (
	"github.com/llzzrrr1997/bamboo/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	// 等待10s才结束执行
	time.Sleep(3 * time.Second)
	_ = c.ISetOkStatus().IJson(map[string]string{"msg": "success"})
}
