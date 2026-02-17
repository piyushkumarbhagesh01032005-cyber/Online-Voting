package routes

import (
	controllers "onlinevoting/controllers"

	"github.com/gorilla/mux"
)

func CandidateRoutes(r *mux.Router) {
	r.HandleFunc("/candidates", controllers.GetCandidates).Methods("GET")
	r.HandleFunc("/candidates", controllers.CreateCandidates).Methods("POST")
	r.HandleFunc("/candidates/{id}", controllers.DeleteCandidates).Methods("DELETE")

}
