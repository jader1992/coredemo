package main

import (
	"gocore/framework"
	"gocore/framework/middleware"
	"log"
	"net/http"
)

func main() {
	// 生成一个新的Handler
	core := framework.NewCore()

	// 注册中间件
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	// 注册路由
	registerRoute(core)

	// 生成server
	server := &http.Server{
		Handler: core,
		Addr: ":8888",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}