package main

import "gocore/framework"

func registerRoute(core *framework.Core)  {
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	subjectApi.Delete("/:id", SubjectDelController)
	subjectApi.Put("/:id", SubjectUpdateController)
	subjectApi.Get("/:id", SubjectGetController)
	subjectApi.Get("/list/all", SubjectListController)

	subjectInnerApi := subjectApi.Group("/info")
	subjectInnerApi.Get("/name", SubjectNameController)
}
