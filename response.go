package context

import (
    "errors"
)

type ResponseBuilder struct {
    ctx             *Context
    status          int
    content         Msg
}

func (ctx Context) BuildResponse(status int) *ResponseBuilder {
    return &ResponseBuilder{
        ctx:     &ctx,
        status:  status,
        content: make(Msg),
    }
}

func (response *ResponseBuilder) Entry(label string, value interface{}) *ResponseBuilder {
    response.content[label] = value
    return response
}

func (response *ResponseBuilder) Send() error {
    response.ctx.StatusCode(response.status)
    if _, err := response.ctx.JSON(response.content); err != nil {
        return errors.New("could not build response: " + err.Error())
    }
    return nil
}

