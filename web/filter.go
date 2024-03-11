package web

type filter interface {
	Pre(ctx *Context)
	After(ctx *Context)
}
