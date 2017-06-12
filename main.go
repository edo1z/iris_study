package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/view"
)

func main() {
	app := iris.New()
	app.AttachView(view.HTML("./views", ".html").Reload(true))

	app.Handle("GET", "/", top)
	app.Handle("GET", "/about", func(ctx context.Context) {
		ctx.View("about.html")
	})
	app.Handle("GET", "/users", func(ctx context.Context) {
		ctx.HTML("<p>this is users page.</p>")
	})

	app.StaticWeb("/assets", "./assets")
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}

func top(ctx context.Context) {
	ctx.ViewData("Username", "Taro")
	ctx.View("top.html")
}
