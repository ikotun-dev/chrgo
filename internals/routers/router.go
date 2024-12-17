package routers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ikotun/chrgo/internals/handlers"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	//user routes
	r.Post("/create-user", handlers.CreateUser)

	//session routes
	r.Post("/create-session", handlers.CreateSession)
	r.Get("/ws", handlers.SocketConn)
	return r
}
