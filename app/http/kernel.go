package http

import "github.com/jader1992/gocore/framework/gin"

func NewHttpEngine() (*gin.Engine, error) {
	// 设置gin的模式
	gin.SetMode(gin.ReleaseMode)
	//  默认的engine
	r := gin.Default()

	Routes(r)
	return r, nil
}