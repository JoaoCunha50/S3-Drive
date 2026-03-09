package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func InitSecret(secret []byte) {
	jwtSecret = secret
}

func ExtractToken(c *gin.Context) string {
	tokenBearer := c.GetHeader("Authorization")

	const prefix = "Bearer "
	var token string
	if len(tokenBearer) > len(prefix) && tokenBearer[:len(prefix)] == prefix {
		token = tokenBearer[len(prefix):]
	}

	return token
}

func SignToken(id string, email string, name string, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      id,
		"email":    email,
		"role":     role,
		"name":     name,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func TokenDecode(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret), nil
	})
}
