package web

import "net/http"

type ServerHandler interface {
}

type Server interface {
	Router
	Start(address string) error
}

var _ Server = &HttpServer{}

type HttpServer struct {
	name    string
	handler Handler
	root    Filter
}

/*
Server 处理Http请求  组装Context  交给Filter处理
*/
func (h *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	context := NewContext()
	context.W = writer
	context.R = request
	h.root(context)
}

func NewHttpServer() Server {
	return &HttpServer{}
}

func (h *HttpServer) Route(method string, pattern string, handlerFunc handlerFunc) {
	h.handler.Route(method, pattern, handlerFunc)
}

func (h *HttpServer) Start(address string) error {
	//TODO implement me
	return http.ListenAndServe("/", h)
}
