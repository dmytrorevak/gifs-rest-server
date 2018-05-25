package main

import (
	"fmt"
	"gifs-rest-server/api/web"
	"gifs-rest-server/storage"
	"net/http"
)

func main() {
	db, err := storage.InitDB("gifs_user:Gifs123456@/gifs_db?parseTime=true")
	//var Database = db
	if err != nil {
		fmt.Print(err)
	}

	h := web.GetHandler(*db)
	r := h.GetRouter()
	//fmt.Print(db)

	//d.Migrate(db)
	// gif := storage.Gif{Age: 11}
	// storage.CreateGif(db, &gif)

	// giff := storage.Gif{Age: 12}
	// storage.CreateGif(db, &giff)

	// gifd := storage.Gif{Age: 13}
	// storage.CreateGif(db, &gifd)

	// //storage.GetGif(db)
	// //fmt.Print("cool")
	//fmt.Print(a)
	//g := m.Gif{Age: 7}
	//d.CreateGif(db, &g)
	//d.GetGif(db)
	p := "3000"
	http.ListenAndServe(":"+p, r)

}
