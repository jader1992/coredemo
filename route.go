package main

import (
	"github.com/jader1992/gocore/framework/gin"
	"github.com/jader1992/gocore/framework/middleware"
)

func registerRoute(core *gin.Engine)  {
	// 静态路由+HTTP方法匹配
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())
	subjectApi.DELETE("/:id", SubjectDelController)
	subjectApi.PUT("/:id", SubjectUpdateController)
	subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
	subjectApi.GET("/list/all", SubjectListController)

	subjectInnerApi := subjectApi.Group("/info")
	subjectInnerApi.GET("/name", SubjectNameController)
}
