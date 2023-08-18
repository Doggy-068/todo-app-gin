package middleware

import (
	"net/http"
	"todo-app/constant"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequire(c *gin.Context) {
	apiKey := c.GetHeader("X-API-Key")
	_, err := jwt.Parse(apiKey, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.JwtKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	} else {
		c.Next()
	}
}
