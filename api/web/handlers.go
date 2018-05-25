package web

import (
	"encoding/json"
	"fmt"
	"gifs-rest-server/storage"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	storage gorm.DB
}

func GetHandler(storage gorm.DB) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h Handler) ListGifs(w http.ResponseWriter, r *http.Request) {
	gifsr := storage.GetGifsList(&h.storage)
	// var l = storage.GetGifsList{gifs}
	// fmt.Println(reflect.TypeOf(gifs))
	rs := storage.GifsList{Gifs: gifsr}
	// var rs map[string][]storage.Gif
	// rs["some"] = gifs
	g, err := json.Marshal(rs)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(g)
	w.Write(g)
}


func