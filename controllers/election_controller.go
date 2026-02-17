package controllers

import (
	"encoding/json"
	"net/http"
	"onlinevoting/database"
	"onlinevoting/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetElections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var elections []models.Election

	result := database.DB.Find(&elections)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(elections)

}

func CreateElections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var elections models.Election
	json.NewDecoder(r.Body).Decode(&elections)

	result := database.DB.Create(&elections)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(elections)
}

func DelectElections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid election ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&models.Election{}, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Candidate not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "candidate deleted successfully",
	})

}
