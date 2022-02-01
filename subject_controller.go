package main

import (
	"github.com/llzzrrr1997/bamboo/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson(map[string]string{"msg": "ok, SubjectAddController"})
}

func SubjectListController(c *gin.Context) {
	c.ISetOkStatus().IJson(map[string]string{"msg": "ok, SubjectListController"})
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson(map[string]string{"msg": "ok, SubjectDelController"})
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson(map[string]string{"msg": "ok, SubjectUpdateController"})
}

func SubjectGetController(c *gin.Context) {
	c.ISetOkStatus().IJson(map[string]string{"msg": "ok, SubjectGetController"})
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson(map[string]string{"msg": "ok, SubjectNameController"})
}
