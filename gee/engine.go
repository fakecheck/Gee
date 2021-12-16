package gee

import (
	"net/http"
)

type Engine struct {
	router *Router
}

func New() *Engine {
	engine := &Engine{}
	engine.router = newRouter()
	return engine
}

func (engine *Engine) Get(pattern string, handler func(ctx Context)) {
	engine.router.addRoute("GET", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler func(ctx Context)) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	hd := engine.router.getHandler(req.Method, req.URL.Path)
	hd(ctx)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
