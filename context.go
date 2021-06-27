package context

import (
    "github.com/kataras/iris/v12"
)

type Context struct {
    iris.Context
}

func From(f func(Context)) iris.Handler {
    return func(ctx iris.Context) {
        f(Context{Context: ctx})
    }
}

