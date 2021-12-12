package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]*Tree // all routers
	middlewares []ControllerHandler // 从core这边设置的中间件
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

func (c *Core) Use(middlewares ...ControllerHandler)  {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) Get(url string, handlers ...ControllerHandler)  {
	allHandlers := append(c.middlewares, handlers...)
	err := c.router["GET"].AddRouter(url, allHandlers)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler)  {
	allHandlers := append(c.middlewares, handlers...)
	err := c.router["POST"].AddRouter(url, allHandlers)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler)  {
	allHandlers := append(c.middlewares, handlers...)
	err := c.router["PUT"].AddRouter(url, allHandlers)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler)  {
	allHandlers := append(c.middlewares, handlers...)
	err := c.router["DELETE"].AddRouter(url, allHandlers)
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

// === http method end

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteNodeByRequest(request *http.Request) *node  {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}

// 所有请求都进入这个函数, 这个函数负责路由分发
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request)  {
	ctx := NewContext(request, response)

	// 查找路由
	node := c.FindRouteNodeByRequest(request)
	if node == nil {
		_ = ctx.SetStatus(404).Json("not found")
		return
	}

	ctx.SetHandlers(node.handlers)

	if err := ctx.Next(); err !=nil {
		_ = ctx.SetStatus(500).Json("inner error")
		return
	}
}


