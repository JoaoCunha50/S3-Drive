package router

import (
	"api/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WebRouter(r *gin.RouterGroup, db *gorm.DB) {
    userGroup := r.Group("/users")
    users.RegisterUserRoutes(userGroup, db)
}