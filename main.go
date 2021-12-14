// Copyright 2021 jade1992.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	coreHttp "github.com/jader1992/gocore/app/http"
	"github.com/jader1992/gocore/app/provider/demo"
	"github.com/jader1992/gocore/framework/gin"
	"github.com/jader1992/gocore/framework/middleware"
	"github.com/jader1992/gocore/framework/provider/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

)

func main() {
	// 生成一个新的Handler
	core := gin.New()

	// 绑定具体的服务提供者
	core.Bind(&demo.DemoProvider{}) // 绑定DemoProvide
	core.Bind(&app.GocoreAppProvider{}) // 绑定GocoreAppProvider

	// 注册中间件
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	// 注册路由
	coreHttp.Routes(core)

	// 生成server
	server := &http.Server{
		Handler: core,
		Addr: ":8888",
	}

	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)

	// 因为使用 Ctrl 或者 kill 命令，它们发送的信号是进入 main 函数的，即只有 main 函数所在的 Goroutine 会接收到，
	// 所以必须在 main 函数所在的 Goroutine 监听信号。
	//  关注的监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束，控制最多等待的时间
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// server.Shutdown方法是个阻塞方法，一旦执行之后，他会阻塞当前的goroutine，并且在所有连接请求结束后，才继续往后执行
	// 这种方式与grace优雅关闭服务比较相似
	// 内部实现：是通过双层循环：第一层循环是定时无限循环，每过ticker的间隔时间，就进入第二层循环；第二层循环是遍历连接中的所有请求，
	// 如果已经处理完操作处于 Idle 状态，就关闭连接，直到所有连接都关闭，才返回。
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("server shutdown:", err)
	}
}
