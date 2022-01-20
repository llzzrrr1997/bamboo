package main

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8888",
	}
	server.ListenAndServe()
}
