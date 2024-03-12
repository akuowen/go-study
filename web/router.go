package web

type Router interface {
	Route(method string, pattern string)
}
