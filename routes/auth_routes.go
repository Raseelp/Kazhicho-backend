package routes

import (
	"github.com/gin-gonic/gin"
	"kazhicho-backend/controllers"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
	}
}
