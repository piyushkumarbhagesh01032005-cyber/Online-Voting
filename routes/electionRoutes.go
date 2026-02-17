package routes

import (
	"onlinevoting/controllers"

	"github.com/gorilla/mux"
)

func ElectionRoutes(r *mux.Router) {
	r.HandleFunc("/election", controllers.GetElections).Methods("GET")
	r.HandleFunc("/election", controllers.CreateElections).Methods("POST")
	r.HandleFunc("/electoin/{id}", controllers.DelectElections).Methods("DELETE")
}
