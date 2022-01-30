package main

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"net/http"
)

func FooControllerHandler(c *framework.Context) error {
	c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "hello,world!"})
	return nil
}
