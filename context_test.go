package context

import (
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	// "github.com/stretchr/testify/assert"
)

func testingApp() *iris.Application {
    app := iris.New()

    app.Get("/dummy", For(dummyEndpoint))
    app.Get("/dummy-fail", For(dummyFailEndpoint))

    return app
}

func dummyEndpoint(ctx Context) {
    _ = ctx.BuildResponse(iris.StatusOK).Entry("msg", "all good").Send()
}

func dummyFailEndpoint(ctx Context) {
    err := ctx.BuildResponse(iris.StatusOK).Entry("msg", make(chan int)).Send()
    if err != nil {
        ctx.BuildResponse(iris.StatusInternalServerError).Entry("error", "could not reply").Send()
    }
}

func TestHello(t *testing.T) {

    app := testingApp()
    e := httptest.New(t, app)

    e.GET("/dummy").Expect().
        Status(iris.StatusOK).
        JSON(httpexpect.ContentOpts{MediaType: "application/json"}).
        Object().
            Value("msg").Equal("all good")

    e.GET("/dummy-fail").Expect().
        Status(iris.StatusInternalServerError).
        JSON(httpexpect.ContentOpts{MediaType: "application/json"}).
        Object().
            Value("error").Equal("could not reply")



}

