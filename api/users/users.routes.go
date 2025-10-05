package users

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(r *gin.Engine, repo *UserRepository) {
	handler := NewUserHandler(repo)

	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)
	r.GET("/users/login", handler.LoginUser)
}
