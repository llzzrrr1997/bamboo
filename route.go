package main

import "github.com/llzzrrr1997/bamboo/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
