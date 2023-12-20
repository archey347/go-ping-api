package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return r
}
