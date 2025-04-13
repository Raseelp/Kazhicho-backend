package routes

import (
	"github.com/gin-gonic/gin"
	"kazhicho-backend/middleware"
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
	spot := r.Group("/foodspot")
	spot.Use(middleware.AuthMiddleware())
	{
		spot.POST("/request-foodspot", services.RequestRegisterFoodSpots)
	}
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.GET("/get-foodspot-requests", services.GetfoodSpotRequests)
	}
}
