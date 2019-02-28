package main

import (
    "github.com/kataras/iris"
)

var app *iris.Application

func main() {
    app = iris.Default()

    app.RegisterView(iris.HTML("./templates", ".html")) // 注册 view 目录
    app.StaticWeb("static", "./static")     // 初始化静态页面目录

    app.Get("/index", func(ctx iris.Context) {
        ctx.ViewData("name", "wxnacy")
        ctx.View("index.html")
    })

    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8080"))
}
