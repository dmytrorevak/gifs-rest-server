package main

import (
	"fmt"
	"gifs-rest-server/api/web"
	"gifs-rest-server/storage"
	"net/http"
)

func main() {
	db, err := storage.InitDB("gifs_user:Gifs123456@/gifs_db?parseTime=true")
	if err != nil {
		fmt.Printf("Error main creation db: %v", err)
	}

	handler := web.NewHandler(db)
	router := handler.GetRouter()
	fmt.Println("Listen at :3000")
	http.ListenAndServe(":3000", router)
}
