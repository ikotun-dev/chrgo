package main

import (
	"net/http"

	"github.com/ikotun/chrgo/internals/config"
	"github.com/ikotun/chrgo/internals/routers"
	log "github.com/sirupsen/logrus"
)

func main() {

	println("Hello World")
	config.InitDB()
	router := routers.InitRouter()

	log.Info("Server is starting on port 8000")

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Error("Error starting server: ", err)

	}
}
