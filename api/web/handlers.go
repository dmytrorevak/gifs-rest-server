package web

import (
    "net/http"
    //"gifs-rest-server/storage"
    "fmt"
	"gifs-rest-server/storage"
)


func listGifs(w http.ResponseWriter, r *http.Request) {
    gifs := storage.GetGifsList(Database)
    fmt.Print(gifs)
}
