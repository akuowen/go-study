package web

type ServerHandler interface {
}

type Server interface {
	Router
}

type HttpServer struct {
	name string
}

func (h HttpServer) Route(method string, pattern string) {
	//TODO implement me
	panic("implement me")
}

//URL --> FILTER --> HANDLER
