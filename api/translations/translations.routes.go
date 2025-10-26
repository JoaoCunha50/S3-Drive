package translations

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTranslationRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewTranslationRepository(db)
	handler := NewTranslationHandler(repo)
	
	r.GET("/", handler.GetTranslations)
	r.POST("/", handler.CreateTranslation)
}