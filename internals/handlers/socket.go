package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ikotun/chrgo/internals/config"
	"github.com/ikotun/chrgo/internals/models"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type message struct {
	Content string `json:"content"`
	User    string `json:"user"`
}

// TODO: extract session id in the url
// TODO: create a new message and save it in the database

func createSession(socket_id string) (uint, error) {
	newSession := models.Session{
		SocketId: socket_id,
	}
	result := config.DB.Create(&newSession)
	log.Info("Session affected : ", result.RowsAffected)
	if result.Error != nil {
		return 0, result.Error
	}
	return newSession.ID, nil

}

func SocketConn(w http.ResponseWriter, r *http.Request) {

	socketID := r.URL.Query().Get("socket_id")

	createdSessionID, err := createSession(socketID)

	if err != nil {
		log.Error("Could not create session: ", err)
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		panic(err)
	}

	for {
		messageType, p, err := conn.ReadMessage()
		log.Info("Message received: ", string(p))
		log.Info("Message type: ", messageType)

		// create a new message
		var msg message
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Error("Could not unmarshal message: ", err)
		}
		log.Info("Message content: ", msg.Content)
		//create message
		newMessage := models.Message{
			Text:      msg.Content,
			UserType:  msg.User,
			SessionID: createdSessionID,
		}

		result := config.DB.Create(&newMessage)

		log.Info("Message created: ", result.Error)
		log.Info("Message affected : ", result.RowsAffected)

		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

}
