package routes

import (
	"github.com/dustin-ward/recordrank/controllers"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Misc Routes
	router.HandleFunc("/", controllers.HomePagePost).Methods("POST")
	router.HandleFunc("/", controllers.HomePageGet).Methods("GET")

	// Auth Routes
	router.HandleFunc("/auth/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")

	//

	return router
}
