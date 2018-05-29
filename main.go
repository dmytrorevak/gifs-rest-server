package main

import (
	"fmt"
	"github.com/dmytrorevak/gifs-rest-server/api/web"
	"github.com/dmytrorevak/gifs-rest-server/storage"
	"log"
	"net/http"
)

func main() {
	db, err := storage.InitDB("gifs_user:Gifs123456@/gifs_db?parseTime=true")
	if err != nil {
		fmt.Printf("Error main creation db: %v", err)
		return
	}
	defer db.Close()

	handler := web.NewHandler(db)
	router := handler.GetRouter()
	fmt.Println("Listen at :3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
