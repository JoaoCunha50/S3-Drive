package router

import (
	"api/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupMainRouter(r *gin.Engine, db *gorm.DB) {
	webGroup := r.Group("/web")
	WebRouter(webGroup, db)
}
