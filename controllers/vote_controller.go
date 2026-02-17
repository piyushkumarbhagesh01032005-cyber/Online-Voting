package controllers

import (
	"encoding/json"
	"net/http"
	"onlinevoting/database"
	"onlinevoting/models"
)

func GetVotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "Application/json")

	var votes []models.Vote

	result := database.DB.Find(&votes)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(votes)

}

func CreateVotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content- Type", "Application/json")

	var votes models.Vote

	json.NewDecoder(r.Body).Decode(&votes)

	result := database.DB.Create(&votes)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(votes)

}
