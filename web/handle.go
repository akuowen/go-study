package web

type handlerFunc func(c *Context)

type handler interface {
	Router

	ServerHttp(ctx *Context)
}
