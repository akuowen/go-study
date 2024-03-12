package web

type handlerFunc func(c *Context)

// Router 路由  将url添加到映射上
type Router interface {
	Route(method string, pattern string, handlerFunc handlerFunc)
}
