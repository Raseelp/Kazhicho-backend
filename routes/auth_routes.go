package routes

import (
	"github.com/gin-gonic/gin"
	"kazhicho-backend/services"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", services.Register)
	}
	{
		auth.POST("/login", services.Login)
	}
}
