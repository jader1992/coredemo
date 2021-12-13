package main

import (
	"github.com/jader1992/gocore/framework/gin"
	"github.com/jader1992/gocore/provider/demo"
)

func SubjectAddController(c *gin.Context) {
	_ = c.ISetOkStatus().IJson("ok, SubjectAddController")
}

func SubjectListController(c *gin.Context) {
	// 获取demo服务实例
	demoService := c.MustMake(demo.KEY).(demo.Service)

	// 调用服务实例的方法
	foo := demoService.GetFoo()

	_ = c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context)  {
	_ = c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	_ = c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectGetController")
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}

