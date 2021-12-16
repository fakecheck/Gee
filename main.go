package main

import (
	"log"
	"net/http"
	"wangyankai/gee/gee"
)

func main() {
	Gee := gee.New()
	Gee.Get("/", indexHandler)
	Gee.Get("/hello", helloHandler)
	Gee.Post("/login", loginHandler)
	log.Fatal(Gee.Run(":9999"))
}

func indexHandler(ctx gee.Context) {
	ctx.SetHtml(http.StatusOK, "<h1>Hello Gee</h1>")
}

func helloHandler(ctx gee.Context) {
	ctx.SetStringResp(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
}

func loginHandler(ctx gee.Context) {
	ctx.SetJsonResp(http.StatusOK, gee.Header{
		"username": ctx.PostForm("username"),
		"password": ctx.PostForm("password"),
	})
}
