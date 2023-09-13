package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

/*
	Contains middlewares
	Service discovery
	Logging
	Monitoring

	Suitable for production microservices

	Supports call using gRPC, HTTP and
	anyother protocol

	For large-scale microservices

	Distributed systems

	High-performance microservices

	API gateways

	Large-scale production apps that scale
	to million of users

	Circuit breaking
*/

func main() {
	// Initialize the logger middleware
	logger := middleware.Chain(
		middleware.Recovery(),
		middleware.Trace(),
		middleware.Log(log.DefaultLogger),
	)

	// Create a service with the logger middleware
	app := kratos.New(
		kratos.Name("my-service"),
		kratos.Server(
			http.NewServer(":8000", logger),
		),
	)
	app.Run()
}