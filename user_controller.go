package main

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"net/http"
	"time"
)

func UserLoginController(c *framework.Context) error {
	// 等待10s才结束执行
	time.Sleep(3 * time.Second)
	_ = c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "success"})
	return nil
}
