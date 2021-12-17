package http

import (
	"github.com/jader1992/gocore/app/http/module/demo"
	"github.com/jader1992/gocore/framework/gin"
	"github.com/jader1992/gocore/framework/middleware"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	// 设置静态路径
	r.Static("/dist/", "./dist/")

	// 使用链路追踪中间件
	r.Use(middleware.Trace())
	demo.Register(r) // 注册路由
}
