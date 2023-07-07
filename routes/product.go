package routes

import (
	"go-native-jwt/controllers"
	"go-native-jwt/middleware"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	router := r.PathPrefix("/products").Subrouter()
	router.Use(middleware.Auth)

	router.HandleFunc("/", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteProduct).Methods("DELETE")
}