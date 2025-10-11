package users

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(r *gin.Engine, repo *UserRepository) {
	handler := NewUserHandler(repo)

	r.POST("/", handler.CreateUser)
	r.GET("/:id", handler.GetUser)
	r.GET("/login", handler.LoginUser)
}
