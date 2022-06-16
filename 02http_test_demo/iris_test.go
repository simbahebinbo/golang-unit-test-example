package http_test_demo

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/kataras/iris/v12"
	"net/http"
	"testing"
)

func TestInfoHandler(t *testing.T) {
	e := irisTester(t)
	t1 := e.GET("/info").
		Expect().
		Status(http.StatusOK)
	t1.Body().Equal("info")
}

func IrisHandler() http.Handler {
	app := iris.New()
	RegisterHttpHandler(app)

	if err := app.Build(); err != nil {
		app.Logger().Error(err)
	}
	return app
}

func irisTester(t *testing.T) *httpexpect.Expect {
	handler := IrisHandler()

	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}
