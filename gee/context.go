package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Header map[string]interface{}

type Context struct {
	Req *http.Request
	w   http.ResponseWriter

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) Context {
	return Context{
		Req:    req,
		w:      w,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (ctx Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx Context) SetStatus(status int) {
	ctx.StatusCode = status
	ctx.w.WriteHeader(status)
}

func (ctx Context) SetHeader(k, v string) {
	ctx.w.Header().Set(k, v)
}

func (ctx Context) SetStringResp(status int, format string, values ...interface{}) {
	ctx.SetHeader("Context-Type", "text/plain")
	ctx.SetStatus(status)
	_, err := ctx.w.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		panic(fmt.Sprintf("[Gee] context.SetStringResp Write error:%+v", err))
	}
}

func (ctx Context) SetJsonResp(code int, obj interface{}) {
	ctx.SetHeader("Context-Type", "text/plain")
	ctx.SetStatus(code)
	encoder := json.NewEncoder(ctx.w)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.w, err.Error(), 500)
	}
}

func (ctx Context) SetData(code int, data []byte) {
	ctx.SetStatus(code)
	_, err := ctx.w.Write(data)
	if err != nil {
		panic(fmt.Sprintf("[Gee] Context.SetData wirte failed, err:%+v ", err))
	}
}

func (ctx Context) SetHtml(code int, html string) {
	ctx.SetStatus(code)
	ctx.SetHeader("ContextType", "text/html")
	_, err := ctx.w.Write([]byte(html))
	if err != nil {
		panic(fmt.Sprintf("[Gee] Context.SetHtml wirte failed, err:%+v ", err))
	}
}
