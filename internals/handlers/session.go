package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ikotun/chrgo/internals/models"
	"github.com/ikotun/chrgo/internals/responses"
)

// type session struct {
// 	SocketId string  `json:"socket_id"`
// 	ThreadID *string `json:"Thread_id"`
// }

func CreateSession(w http.ResponseWriter, r *http.Request) {

	var session models.Session
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		responses.RequestError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newSession := models.Session{
		SocketID:  session.SocketID,
		ChatbotID: session.ChatbotID,
	}
	log.Println(newSession)
	responses.RequestSuccess(w, "Session created successfully", http.StatusCreated)
}
