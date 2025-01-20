package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deepanshumishraa/models"
	"gorm.io/gorm"
)

type APIKeyRequest struct {
	APIKEY string `json:"apikey"`
}

func GetUserByAPIKey(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req APIKeyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.APIKEY == "" {
			http.Error(w, "API Key is required", http.StatusBadRequest)
			return
		}

		var user models.User
		result := db.Where("api_key = ?", req.APIKEY).First(&user)
		if result.Error != nil {
			log.Printf("Error finding user: %v", result.Error)
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
