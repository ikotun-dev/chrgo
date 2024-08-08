package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ikotun/chrgo/internals/config"
	"github.com/ikotun/chrgo/internals/models"
)

type createUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// create a user
	var user createUser

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{"error": "Invalid request"}
		json.NewEncoder(w).Encode(errorMessage)

		return
	}

	newUser := models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	result := config.DB.Create(&newUser)
	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := map[string]string{"error": "Could not create user"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User Created")

}
