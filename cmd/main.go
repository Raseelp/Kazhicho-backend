package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kazhicho-backend/routes"
	"kazhicho-backend/services"
	"log"
	"net/http"

	"kazhicho-backend/config"
)

func main() {
	//Connecting DB
	config.InitConfig()
	services.InitCollections(config.DB)
	//initializing Routes
	router := gin.Default()
	routes.AuthRoutes(router)
	routes.AdminRoutes(router)
	routes.UserAndFoodSpotsRoutes(router)
	router.Run(":8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Kazhicho? backend is running")
	})
	fmt.Println("Server is starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
