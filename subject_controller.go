package main

import "github.com/jader1992/gocore/framework/gin"

func SubjectAddController(c *gin.Context) {
	_ = c.ISetOkStatus().IJson("ok, SubjectAddController")
}

func SubjectListController(c *gin.Context) {
	_ = c.ISetOkStatus().IJson("ok, SubjectListController")
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

