package routes

import (
	controllers "onlinevoting/controllers"

	"github.com/gorilla/mux"
)

func VoteRoutes(r *mux.Router) {
	r.HandleFunc("/vote", controllers.GetVotes).Methods("GET")
	r.HandleFunc("/vote", controllers.CreateVotes).Methods("POST")

}
