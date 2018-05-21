package main

import (
    "fmt"
    "net/http"
    "gifs-rest-server/storage"
    "gifs-rest-server/api/web"
)




func main() {
    db, err := storage.InitDB("gifs_user:Gifs123456@/gifs_db")
    if err != nil {
        fmt.Printf("Error while init db: %v", err)
        return
    }
    //d.Migrate(db)
    // gif := models.Gif{Age:11, }
    // database.CreateGif(db, &gif)

    storage.GetGif(db)
    fmt.Print("cool")
    //fmt.Print(a)
    // g := m.Gif{Age: 7}
    // d.CreateGif(db, &g)
    //d.GetGif(db)
    r := web.GetRouter()

	http.ListenAndServe(":3000", r)
}
