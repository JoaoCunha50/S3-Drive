package main

import (
	"api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env := config.LoadConfig()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":" + env.PORT)
}