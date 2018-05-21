package web

import (
	"net/http"
	"github.com/go-chi/chi"
	"gifs-rest-server/storage"
	"fmt"
)

var Database, err = storage.InitDB("gifs_user:Gifs123456@/gifs_db")

func Serve(port string, router chi.Router)  {
	//db, err := storage.InitDB("gifs_user:Gifs123456@/gifs_db")
	//var Database = db
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Print(db)
	http.ListenAndServe(":" + port, router)

}
