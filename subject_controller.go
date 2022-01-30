package main

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"net/http"
)

func SubjectAddController(c *framework.Context) error {
	_ = c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "ok, SubjectAddController"})
	return nil
}

func SubjectListController(c *framework.Context) error {
	_ = c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "ok, SubjectListController"})
	return nil
}

func SubjectDelController(c *framework.Context) error {
	_ = c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "ok, SubjectDelController"})
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	_ = c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "ok, SubjectUpdateController"})
	return nil
}

func SubjectGetController(c *framework.Context) error {
	_ = c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "ok, SubjectGetController"})
	return nil
}

func SubjectNameController(c *framework.Context) error {
	_ = c.SetStatus(http.StatusOK).Json(map[string]string{"msg": "ok, SubjectNameController"})
	return nil
}
