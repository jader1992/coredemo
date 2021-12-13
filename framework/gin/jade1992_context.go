package gin

import (
	"context"
	"github.com/jader1992/gocore/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// Bind 实现了engine的绑定
func (engine *Engine) Bind(provide framework.ServiceProvider) error {
	return engine.container.Bind(provide)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// context 实现了container的几个封装

// Make 实现了context的Make封装
func (ctx *Context) Make(key string) (interface{}, error)  {
	return ctx.container.Make(key)
}

func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
