package framework

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]*Tree // all routers
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

// === http method wrap

func (c *Core) Get(url string, handler ControllerHandler)  {
	err := c.router["GET"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler)  {
	err := c.router["POST"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler)  {
	err := c.router["PUT"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler)  {
	err := c.router["DELETE"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

// === http method end

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler  {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// 所有请求都进入这个函数, 这个函数负责路由分发
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request)  {
	ctx := NewContext(request, response)

	// 查找路由
	router := c.FindRouteByRequest(request)
	if router == nil {
		_ = ctx.Json(404, "not found")
		return
	}

	if err := router(ctx); err !=nil {
		fmt.Println(2)
		_ = ctx.Json(500, "inner error")
		return
	}
}


