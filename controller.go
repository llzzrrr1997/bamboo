package main

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"net/http"
)

func FooControllerHandler(c *framework.Context) error {
	c.Json(http.StatusOK, map[string]string{"msg": "hello,world!"})
	return nil
}
