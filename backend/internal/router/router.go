package router

import (
	"github.com/JoaoCunha50/S3-Drive/backend/internal/info"
	"github.com/JoaoCunha50/S3-Drive/backend/internal/translation"
	"github.com/JoaoCunha50/S3-Drive/backend/internal/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupMainRouter(r *gin.Engine, db *mongo.Database) {
	webGroup := r.Group("/web")

	userRepo := user.NewRepository(db)
	translationRepo := translation.NewRepository(db)
	infoService := info.NewService(userRepo, translationRepo)

	userHandler := user.NewHandler(userRepo)
	translationHandler := translation.NewHandler(translationRepo)

	usersGroup := webGroup.Group("/users")
	usersGroup.POST("/login", userHandler.LoginUser)
	usersGroup.GET("/:id", userHandler.GetUser)
	usersGroup.POST("", userHandler.CreateUser)

	translationsGroup := webGroup.Group("/translations")
	translationsGroup.GET("", translationHandler.GetTranslations)
	translationsGroup.POST("", translationHandler.CreateTranslation)

	infoGroup := webGroup.Group("/info")
	info.RegisterRoutes(infoGroup, infoService)
}
