package router

import (
	"api/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    userGroup := r.Group("/users")
    users.RegisterUserRoutes(userGroup, db)
}