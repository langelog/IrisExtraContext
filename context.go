package context

import (
    "github.com/kataras/iris/v12"
)

type Context struct {
    iris.Context
}

func For(f func(Context)) iris.Handler {
    return func(ctx iris.Context) {
        f(Context{Context: ctx})
    }
}

func (ctx Context) SetUser(user interface{}) {
    ctx.Values().Set("user", user)
}

func (ctx Context) GetUser() interface{} {
    return ctx.Values().Get("user")
}

