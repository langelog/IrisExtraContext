package context

import (
    "github.com/kataras/iris/v12"
    "encoding/json"
    "errors"
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

func (ctx Context) ParseBody(targetStructure interface{}) error {
    if err := json.NewDecoder(ctx.Request().Body).Decode(targetStructure); err != nil {
        return errors.New("could not parse input")
    }
    return nil
}

