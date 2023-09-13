package main

import (
	"github.com/gin-gonic/gin"
)

/*
	Is a light and flexible framework for
	high performance APIs

	Handles a high volume of traffic

	Open source

	It's considered and advanced framework for
	its stable features

	Comprehensible documentation

	Scalable

	Production-level apps

	Ease of maintainance

	Larger than Echo and Fiber

	Not for small projects
*/

func main(){
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	r.Run(":8080")
}