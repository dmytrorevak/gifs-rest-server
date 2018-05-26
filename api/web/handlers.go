package web

import (
	"encoding/json"
	"fmt"
	"gifs-rest-server/storage"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/go-chi/chi"
	"strconv"
	"io/ioutil"
)

type Handler struct {
	storage gorm.DB
}

func NewHandler(storage *gorm.DB) *Handler {
	return &Handler {
		storage: *storage,
	}
}

func (h Handler) GetGifs(w http.ResponseWriter, r *http.Request) {
	gifs := storage.GetGifsList(&h.storage)
	resp, err := json.Marshal(gifs)
	if err != nil {
		fmt.Printf("Error serialize gifs list: %v", err)
	}

	w.Write(resp)
}

func (h Handler) GetGif(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "gifID")
	id, err := strconv.Atoi(strID)
	if err != nil {
		fmt.Printf("Error convert gifId: %v", err)
		w.WriteHeader(400)
	}

	g := storage.GetGif(&h.storage, id)
	resp, err := json.Marshal(g)
	if err != nil {
		fmt.Printf("Error serialize gif: %v", err)
		w.WriteHeader(404)
	}

	w.Write(resp)
}

func (h Handler) CreateGif(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error parse body on POST: %v", err)
		w.WriteHeader(400)
	}

	var g storage.Gif
	json.Unmarshal(body, &g)
	storage.CreateGif(&h.storage, &g)
	w.WriteHeader(201)
}

func (h Handler) UpdateGif(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "gifID")
	id, err := strconv.Atoi(strID)
	if err != nil {
		fmt.Printf("Error convert gifId: %v", err)
		w.WriteHeader(400)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error parse body on POST: %v", err)
		w.WriteHeader(400)
	}

	var data storage.Gif
	json.Unmarshal(body, &data)
	storage.UpdateGif(&h.storage, id, &data)
	w.WriteHeader(204)
}

func (h Handler) DeleteGif(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "gifID")
	id, err := strconv.Atoi(strID)
	if err != nil {
		fmt.Printf("Error convert gifId: %v", err)
		w.WriteHeader(400)
	}

	storage.DeleteGif(&h.storage, id)
	w.WriteHeader(204)
}
