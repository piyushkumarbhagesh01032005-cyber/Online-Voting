package controllers

import (
	"encoding/json"
	"net/http"
	"onlinevoting/database"
	"onlinevoting/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var candidates []models.Candidate

	result := database.DB.Find(&candidates)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(candidates)

}

func CreateCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var candidates models.Candidate
	json.NewDecoder(r.Body).Decode(&candidates)

	result := database.DB.Create(&candidates)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(candidates)

}

func DeleteCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid candidate ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&models.Candidate{}, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Candidate not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "candidate deleted successfully",
	})

}
