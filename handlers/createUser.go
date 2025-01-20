package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deepanshumishraa/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		apiKey := uuid.NewString()

		user := models.User{
			Name:   req.Name,
			APIKEY: apiKey,
		}

		if err := db.Create(&user).Error; err != nil {
			log.Printf("Error creating user: %v", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}


		log.Printf("User created: %+v", user)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"apiKey": user.APIKEY,
		})
	}
}
