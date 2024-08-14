package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ikotun/chrgo/internals/config"
	"github.com/ikotun/chrgo/internals/models"
	"github.com/ikotun/chrgo/internals/responses"
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
		responses.RequestError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newUser := models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	//hash

	result := config.DB.Create(&newUser)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "23505") {
			responses.RequestError(w, "Conflict! User with this credentials already exist.", http.StatusConflict)
			return
		}
		responses.RequestError(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	//OTP stuff

}
