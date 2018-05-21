package web

import (
    "net/http"
    "github.com/go-chi/chi"
    "gifs-rest-server/api/web"
)


func GetRouter() chi.Router {
    r := chi.NewRouter()
    r.Route("/gifs", func(r chi.Router) {
        r.Get("/", listGifs)
    })

    return r
}

