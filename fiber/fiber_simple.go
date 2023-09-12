package main

import "github.com/gofiber/fiber/v2"

/*
	Inspired from Express Js,
	implements the same way of working
	but for Golang

	Built on top of FastHttp

	It can serve static files like
	HTML, CSS and JS easily by simply
	defining static routes

	Serve the content of multiple directories
	on the same route

	Perfect for building API servers

	Rate limiters

	Low memory footprint

	Perfect for startups. You can build MVPs
	and scalable systems

	Perfect for chat, video, and video
	streaming servers
*/
func main(){
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":3000")
}