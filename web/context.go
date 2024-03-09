package web

import "net/http"

type context struct {
	W http.ResponseWriter
	R http.Request
}

func (* context) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	//
}