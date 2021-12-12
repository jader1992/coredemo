package main

import "gocore/framework"

func registerRoute(core *framework.Core)  {
	core.Get("foo", FooControllerHandler)
}
