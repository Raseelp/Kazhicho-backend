package main

import (
	"fmt"
	"log"
	"net/http"

	"kazhicho-backend/config"
)

func main() {
	config.InitConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Kazhich? backend is running")
	})
	fmt.Println("Server is starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
