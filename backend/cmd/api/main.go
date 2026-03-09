package main

import (
	"log"

	"github.com/JoaoCunha50/S3-Drive/backend/internal/config"
	"github.com/JoaoCunha50/S3-Drive/backend/internal/infra"
	"github.com/JoaoCunha50/S3-Drive/backend/internal/router"
	"github.com/JoaoCunha50/S3-Drive/backend/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Env
	defer pkg.InitLogger("dev")()

	client, err := infra.NewMongoClient(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	db := infra.GetDatabase(client, cfg.MongoDatabase)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	r.OPTIONS("/*path", func(c *gin.Context) { c.Status(204) })

	router.SetupMainRouter(r, db)
	r.Run(":" + cfg.Port)
}
