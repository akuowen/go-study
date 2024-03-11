package web

import "net/http"

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

//func (* context) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
//	//
//}

func  NewContext() *Context {
	return &Context{}
}