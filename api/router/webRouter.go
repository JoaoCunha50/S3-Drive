package router

import (
	"api/translations"
	"api/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WebRouter(r *gin.RouterGroup, db *gorm.DB) {
    userGroup := r.Group("/users")
	translationsGroup := r.Group("/translations")

    users.RegisterUserRoutes(userGroup, db)
	translations.RegisterTranslationRoutes(translationsGroup, db)
}