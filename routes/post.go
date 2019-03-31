package routes

import (
	"github.com/bennyzanuar/goblog/controllers"

	"github.com/gorilla/mux"
)

// InitPostRoutes - Create routes for CRUD
func InitPostRoutes(r *mux.Router) {
	s := r.PathPrefix("/posts").Subrouter()

	s.HandleFunc("/", controllers.GetAllPost).Methods("GET")
}
