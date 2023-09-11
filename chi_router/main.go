package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/*
	Quick and light framework
	Build demos and proof of concept
	Modular Go Framework
	Handling millions of web requests
	No external dependnecies
	Uses net/http

	Middlewares
		Logger
		Recoverer
		Throttler

	Focus on writing business logic
	Any net/http middleware also compatible

	REST API services thar are kept maintainable
	as your project grows and changes

	Focuses on
		Developer productivity
		Project structure
		Maintainability
		Simplicity

	Great for building MVPs
	Short learning curve
	Anything that requires fast iterations
	Suitable for Hackathons and internal
	products
*/
func main(){
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}