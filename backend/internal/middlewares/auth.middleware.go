package middlewares

import (
	"github.com/JoaoCunha50/S3-Drive/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HasAuthentication(c *gin.Context) {
	tokenStr := utils.ExtractToken(c)
	if tokenStr == "" {
		c.Next()
		return
	}

	token, err := utils.TokenDecode(tokenStr)
	if err != nil || !token.Valid {
		c.Next()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Next()
		return
	}

	currentUser := map[string]any{
		"id":       claims["sub"].(string),
		"email":    claims["email"].(string),
		"username": claims["username"].(string),
		"name":     claims["name"].(string),
		"role":     claims["role"].(string),
	}

	c.Set("currentUser", currentUser)
	c.Next()
}
