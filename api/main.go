package main

import (
	"api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env := config.LoadConfig()
	db := config.DBconnection(env.DATABASE_URL)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	
	r.Run(":" + env.PORT)
}