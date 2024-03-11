package web

import "net/http"

type ServerHandler interface {
	
}

type Server struct {
	W http.ResponseWriter
	R *http.Request
	handler ServerHandler
}


//URL --> FILTER --> HANDLER