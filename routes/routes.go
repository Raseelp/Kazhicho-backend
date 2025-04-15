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
		admin.POST("/foodspot-request/:id/approve", services.ApproveFoodSpotRequest)
		admin.DELETE("foodspot-request/:id/reject", services.RejectFoodSpotRequest)
		admin.DELETE("delete-foodspot", services.DeleteFoodSpotAdmin)
	}
}

func UserAndFoodSpotsRoutes(r *gin.Engine) {
	spot := r.Group("/foodspot")
	spot.Use(middleware.AuthMiddleware())
	{
		spot.POST("/request-foodspot", services.RequestRegisterFoodSpots)
		spot.POST("/:foodspot_id/add-fooditem", services.AddFoodItemToFoodSpot)
		spot.POST("/:foodspot_id/add-Deal", services.AddDeal)
		spot.PUT("/edit-fooditem", services.EditFoodItem)
		spot.PUT("/remove-fooditem-inmenu", services.RemoveFoodItemFromMenu)
		spot.DELETE("/delete-deal", services.DeleteDeal)
	}
	user := r.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.POST("/:foodspot_id/add-review", services.AddReview)
		user.POST("/add-list", services.AddList)
		user.POST("/upload-reel", services.UploadReel)
	}
}
