package web

import (
    "github.com/go-chi/chi"
)


func GetRouter() chi.Router {
    r := chi.NewRouter()
    r.Route("/gifs", func(r chi.Router) {
        r.Get("/", listGifs)
    })

    return r
}

