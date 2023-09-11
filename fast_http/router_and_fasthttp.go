package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

/*
	Provides a web server gateway interface
	and HTTP server engine drivers

	It is not a complete framework, so
	Frameworks can be written on top of it

	It's been designed with a very specific
	use case in mind: handling thousands of
	small to medium requests per second
	with a consistently low response time

	Limited to specific use cases, for
	example:
		Chat servers
		Chat rooms
		Video servers

	Supports web sockets

	It's 10 times faster than net/http

	It's recommended to write all the handle
	function requests from scratch in order to
	user all of the advantages of the framework

	Natively doesn't have a router

 	Popular among beginners

	Can reuse existing request objects and
	optimize by reducing overhead

	Returns a string of bytes rather than
	a string, using less memory allocation
	and copy

	Optimized for low memory usage

	Projects not requiring a heavy
	business layer

	Not the best for
		CRM
		Project Management tools
		ERP

	Perfect for light clients
	Building layers on top of existing APIs
	Speed
	Handle concurrent connections
	Making requests to other systems
	Can compose data for usage and presentation
*/
func Index(ctx *fasthttp.RequestCtx){
	ctx.WriteString("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx){
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}