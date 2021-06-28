package context

import (
	"testing"
    "log"
	"github.com/gavv/httpexpect"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
    irisContext "github.com/kataras/iris/v12/context"
	"github.com/stretchr/testify/assert"
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

func TestResponseBuilder(t *testing.T) {

    app := testingApp()
    e := httptest.New(t, app)

    // check for simple message response building
    e.GET("/dummy").Expect().
        Status(iris.StatusOK).
        JSON(httpexpect.ContentOpts{MediaType: "application/json"}).
        Object().
            Value("msg").Equal("all good")

    // check for reception of errors
    e.GET("/dummy-fail").Expect().
        Status(iris.StatusInternalServerError).
        JSON(httpexpect.ContentOpts{MediaType: "application/json"}).
        Object().
            Value("error").Equal("could not reply")

}

// --

type SampleUser struct {
    Name    string
}

func TestUserSetGet(t *testing.T) {
    // simulating context
    app := iris.New()
    ctx := Context{Context: irisContext.NewContext(app)}

    // store user in context
    userInput := &SampleUser{Name: "Peter"}
    ctx.SetUser(userInput)

    // retrieve user from context
    user, ok := ctx.GetUser().(*SampleUser)
    if !ok {
        log.Println("NOK")
    }
    
    // check
    assert.True(t, ok)
    assert.Equal(t, "Peter", user.Name)
}

