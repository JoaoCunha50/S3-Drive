package main

import (
	"api/config"
	"api/database"
	"api/router"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadConfig()

	db := database.DBconnection(env.DATABASE_URL)
	err := database.CreateAdmin(env.ADMIN_USERNAME, env.ADMIN_EMAIL, env.ADMIN_PASSWORD, db)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Admin created successfully!")
	}
	
	r := gin.Default()

	router.SetupMainRouter(r, db)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	r.OPTIONS("/*path", func(c *gin.Context) { c.Status(204) })
	
	r.Run(":" + env.PORT)
}