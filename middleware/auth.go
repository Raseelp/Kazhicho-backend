package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"kazhicho-backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//check missing header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization Header is missing"})
			c.Abort()
			return
		}

		//check token format

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Token Format"})
			c.Abort()
			return
		}

		//validate token
		token, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid or Expired Token"})
			c.Abort()
			return
		}

		if Claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", Claims["username"])
			c.Set("exp", Claims["exp"])
		}
		c.Next()

	}
}
