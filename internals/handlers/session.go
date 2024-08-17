package handlers

type session struct {
	SocketId string  `json:"socket_id"`
	ThreadID *string `json:"Thread_id"`
}

func CreateSession() {

	var session session
}
