package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[from docker] ping!",
		})
	})

	router.Run(":4000") // run on port 4000
}
