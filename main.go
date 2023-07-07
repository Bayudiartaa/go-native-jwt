package main

import (
	"go-native-jwt/configs"
	"go-native-jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()
    routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.ProductRoutes(router)
	log.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", router)
}