package main

import "go-study/web"

func main() {
	server := web.NewHttpServer()
	err := server.Start("8080")
	server.Route("GET", "/", func(c *web.Context) {

	})
	if err != nil {
		panic(err)
	}
}
