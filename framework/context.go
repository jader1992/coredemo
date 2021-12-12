package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// Context 自定义context结构
type Context struct {
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context

	handlers []ControllerHandler
	index    int // 当前请求调用链的哪个节点

	// 是否超时标记位
	hasTimeout bool
	// 写保护机制
	writerMux *sync.Mutex

	params map[string]string // url 路由匹配参数
}

// NewContext 初始化上下文

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request: r,
		responseWriter: w,
		ctx: r.Context(),
		writerMux: &sync.Mutex{},
		index: -1,
	}
}

// 注册Context的函数

func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout()  {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// SetHandlers 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler)  {
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string)  {
	ctx.params = params
}

// Next 核心函数，调用context的下一个函数
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

// 获取基础上下文context

func (ctx *Context) BaseContext() context.Context  {
	return ctx.request.Context()
}

// 实现继承的context.Context

func (ctx *Context) Deadline() (deadline time.Time, ok bool)  {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error  {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}