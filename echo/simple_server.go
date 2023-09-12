package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
	High performance, extensible, minimalist
	Go web framework

	Offers robust error handlig

	Lightweight

	Minimalist

	Feature rich

	REST APIs to WebSockets

	File upload and download

	Versatile and general purpose

	Social networks to B2B SaaS

	Great for MVPs

	Scalable

	Community is small relatively
*/

func main(){
	// Echo instance
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}