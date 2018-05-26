package web

import (
	"github.com/go-chi/chi"
)

func (h Handler) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Route("/gifs", func(r chi.Router) {
		r.Get("/", h.GetGifs)
		r.Get("/{gifID}", h.GetGif)
		r.Post("/", h.CreateGif)
		r.Put("/{gifID}", h.UpdateGif)
		r.Delete("/{gifID}", h.DeleteGif)
	})

	return r
}
