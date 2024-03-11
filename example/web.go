package main

import (
	"fmt"
	"go-study/web"
	"net/http"
)

func homapage(w http.ResponseWriter,r *http.Request)  {
	_, err := fmt.Fprintf(w, "this is a homepage")
	if err != nil {
		return 
	}
}

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		context := web.NewContext()
		context.R=request;
		context.W=writer;
	})
	http.HandleFunc("/",homapage)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return 
	}
}