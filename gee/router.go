package gee

import (
	"fmt"
	"net/http"
)

type Router struct {
	routerMap map[string]func(ctx Context)
}

func newRouter() *Router {
	return &Router{routerMap: make(map[string]func(ctx Context))}
}

func (router *Router) addRoute(method, pattern string, handler func(ctx Context)) {
	key := method + "_" + pattern
	h, ok := router.routerMap[key]
	if !ok {
		router.routerMap[key] = handler
		return
	}
	panic(fmt.Sprintf("trying to add multiple handler under Path:%v, existingHandler:%#v, trying to add: %#v", pattern, h, handler))
}

func (router *Router) getHandler(method, path string) func(ctx Context) {
	searchKey := method + "_" + path
	handler, ok := router.routerMap[searchKey]
	if ok {
		return handler
	} else {
		return notFoundHandler
	}
}

func notFoundHandler(ctx Context) {
	ctx.SetStringResp(http.StatusNotFound, "404 NOT FOUND:%s \n", ctx.Path)
}
