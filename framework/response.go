package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"text/template"
)

type IResponse interface {
	// json 输出
	Json(obj interface{}) IResponse // 返回 IResponse，这样能允许使用方进行链式调用

	// Jsonp输出
	Jsonp(obj interface{}) IResponse

	// xml输出
	Xml(obj interface{}) IResponse

	// html输出
	Html(template string, obj interface{}) IResponse

	// string输出
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httponly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置200状态
	SetOkStatus() IResponse
}

/********************response 状态相关*****************************/

func (ctx *Context) SetStatus(code int) IResponse  {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}

func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool,
httpOnly bool) IResponse{
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name: key,
		Value: url.QueryEscape(val),
		MaxAge: maxAge,
		Path: path,
		Domain: domain,
		SameSite: 1,
		Secure: secure,
		HttpOnly: httpOnly,
	})

	return ctx
}

/********************responseWriter 输出相关*****************************/

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}

// 我们知道 HTML 中标签 <script> 中的请求是不受同源策略影响的，那如果能将 B 网站的资源数据，通过 script 标签返回来，
// 是不是就直接解决了跨域问题？确实，JSONP 就是这么设计的，通过 script 标签的源地址，返回数据资源 + JavaScript 代码。

// 使用 ajax 请求的时候，我们是做不到这一点的，因为同源策略限制了网页和数据必须在同一个源下。但是如果 B 网站的接口支持了 JSONP，
// 它能根据请求参数返回一段 JavaScript 代码，类似： callfunc({"id":1, "name": jianfengye})

func (ctx *Context) Jsonp(obj interface{}) IResponse {
	// 获取请求参数callback
	callBackFun, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")

	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callBackFun)

	// 输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}

	// 输出函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}

	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}

	return ctx
}

func (ctx *Context) Xml(obj interface{}) IResponse {
	byte, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/html")
	ctx.responseWriter.Write(byte)
	return ctx
}

func (ctx *Context) Html(file string, obj interface{}) IResponse {
	// 读取模版文件，创建template实例
	t, err := template.New("oupput").ParseFiles(file)
	if err != nil {
		return ctx
	}

	// 执行Execute方法将obj和模版进行结合
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")
	return ctx
}

func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}