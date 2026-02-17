package main

import (
	"fmt"
	"log"
	"net/http"
	"onlinevoting/database"
	"onlinevoting/routes"

	"github.com/gorilla/mux"
)

func main() {

	database.Connect()

	r := mux.NewRouter()

	routes.VoteRoutes(r)
	routes.ElectionRoutes(r)
	routes.CandidateRoutes(r)
	routes.AuthRoutes(r)

	fmt.Println("Server Started at http://localhost:8087")

	log.Fatal(http.ListenAndServe(":8087", r))
}
