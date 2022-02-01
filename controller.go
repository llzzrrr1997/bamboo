package main

import (
	"github.com/llzzrrr1997/bamboo/framework/gin"
	"net/http"
)

func FooControllerHandler(c *gin.Context) {
	c.ISetStatus(http.StatusOK).IJson(map[string]string{"msg": "hello,world!"})
}
