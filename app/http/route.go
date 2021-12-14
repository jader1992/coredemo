package http

import (
	"github.com/jader1992/gocore/app/http/module/demo"
	"github.com/jader1992/gocore/framework/gin"
)

func Routes(r *gin.Engine)  {

	// 设置静态路径
	r.Static("/dist/", "./dist/")

	demo.Register(r) // 注册路由
}
