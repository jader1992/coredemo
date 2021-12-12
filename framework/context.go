package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Context 自定义context结构
type Context struct {
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context
	handle ControllerHandler

	// 是否超时标记位
	hasTimeout bool

	// 写保护机制
	writerMux *sync.Mutex
}

// NewContext 初始化上下文

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request: r,
		responseWriter: w,
		ctx: r.Context(),
		writerMux: &sync.Mutex{},
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

// 获取 query url

func (ctx *Context) QueryAll() map[string][]string{
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string, def int) int  {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		lenVal := len(values)
		if lenVal > 0 {
			intVal, err := strconv.Atoi(values[lenVal-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		lenVal := len(values)
		if lenVal > 0 {
			return values[lenVal-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

// 从 post 获取数据

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		lenVal := len(values)
		if lenVal > 0 {
			intVal, err := strconv.Atoi(values[lenVal-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		lenVal := len(values)
		if lenVal > 0 {
			return values[lenVal-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

// 获取 application/json post

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request == nil {
		return errors.New("ctx.request empty")
	}
	body, err := ioutil.ReadAll(ctx.request.Body)
	if err != nil {
		return err
	}
	ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}
	return nil
}

// 与 response 相关的方法

// 输出json

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	_, _ = ctx.responseWriter.Write(byt)
	return nil
}

// 输出html

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

// 输出text格式

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
