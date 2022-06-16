package http_test_demo

import (
	"github.com/kataras/iris/v12"
)

func InfoHandler(ctx iris.Context) {
	_, _ = ctx.Writef("info")
}

// RegisterHttpHandler Router function handles the iris routers
func RegisterHttpHandler(app *iris.Application) {
	app.Get("/info", InfoHandler)
}
