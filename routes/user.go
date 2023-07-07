package routes

import (
	"go-native-jwt/controllers"
	"go-native-jwt/middleware"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()
	router.Use(middleware.Auth)

	router.HandleFunc("/me", controllers.Me).Methods("GET")
}