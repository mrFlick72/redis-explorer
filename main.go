package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())

	tmpl := iris.HTML("src/views", ".html")
	app.RegisterView(tmpl)
	app.Get("/", index)

	app.Listen(":3000")
}
func index(ctx iris.Context) {
	ctx.View("index.html")
}
