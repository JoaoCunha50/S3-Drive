package main

import (
	"api/config"
	"api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadConfig()
	db := config.DBconnection(env.DATABASE_URL)
	
	r := gin.Default()

	router.SetupRoutes(r, db)
	
	r.Run(":" + env.PORT)
}