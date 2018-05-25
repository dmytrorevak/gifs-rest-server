package web

import (
	"github.com/go-chi/chi"
)

func (h Handler) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Route("/gifs", func(r chi.Router) {
		r.Get("/", h.ListGifs)
	})

	return r
}
