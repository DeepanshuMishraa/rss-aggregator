package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deepanshumishraa/models"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if user.Name == "" || user.APIKEY == "" {
			http.Error(w, "Name and APIKEY are required", http.StatusBadRequest)
			return
		}

		result := db.Create(&user)
		if result.Error != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
