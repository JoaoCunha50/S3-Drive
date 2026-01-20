package router

import (
	"api/data/info"
	"api/data/translations"
	"api/data/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WebRouter(r *gin.RouterGroup, db *gorm.DB) {
    userGroup := r.Group("/users")
	translationsGroup := r.Group("/translations")
	infoGroup := r.Group("/info")

	translationsRepo := translations.NewTranslationRepository(db)
	usersRepo := users.NewUserRepository(db)

    users.RegisterUserRoutes(userGroup, usersRepo)
	translations.RegisterTranslationRoutes(translationsGroup, translationsRepo)
	info.RegisterInfoRoutes(infoGroup, usersRepo, translationsRepo)
}
