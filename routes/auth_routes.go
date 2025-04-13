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
}

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.GET("/get-foodspot-requests", services.GetfoodSpotRequests)
	}
}

func UserAndFoodSpotsRoutes(r *gin.Engine) {
	spot := r.Group("/foodspot")
	spot.Use(middleware.AuthMiddleware())
	{
		spot.POST("/request-foodspot", services.RequestRegisterFoodSpots)
	}
}
