package web

import (
	"encoding/json"
	"fmt"
	"github.com/dmytrorevak/gifs-rest-server/storage"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	//"log"
)

type Handler struct {
	storage gorm.DB
}

func NewHandler(storage *gorm.DB) *Handler {
	return &Handler{
		storage: *storage,
	}
}

func (h Handler) GetGifs(w http.ResponseWriter, r *http.Request) {
	gifs, err := storage.GetGifsList(&h.storage)
	if err != nil {
		//TODO: implement logging + middleware
		return
	}
	resp, err := json.Marshal(gifs)
	if err != nil {
		//TODO: implement logging + middleware
		return
	}

	w.Write(resp)
	// Add status code
}

func (h Handler) GetGif(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "gifID")
	id, err := strconv.Atoi(strID)
	if err != nil {
		//Logging! fmt.Printf("Error convert gifId: %v", err)
		w.WriteHeader(400)
	}

	g, err := storage.GetGif(&h.storage, id)
	if err != nil {
		// ADD logging
		// do not ignore errors!
		return
	}
	resp, err := json.Marshal(g)
	if err != nil {
		//logging fmt.Printf("Error serialize gif: %v", err)
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
	// always check errors
	err = json.Unmarshal(body, &g)
	if err != nil {
		// logging + status code
		return
	}
	err = storage.CreateGif(&h.storage, &g)
	if err != nil {
		// logging + status code
		return
	}
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
	// TODO; validating
	err = json.Unmarshal(body, &data)
	if err != nil {
		//log err
		return
	}
	err = storage.UpdateGif(&h.storage, id, &data)
	if err != nil {
		// logging + status code
		return
	}
	w.WriteHeader(204)
}

func (h Handler) DeleteGif(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "gifID")
	id, err := strconv.Atoi(strID)
	if err != nil {
		//log fmt.Printf("Error convert gifId: %v", err)
		w.WriteHeader(400)
		return
	}

	err = storage.DeleteGif(&h.storage, id)
	if err != nil {
		//logging fmt.Printf("Error convert gifId: %v", err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(204)
}
