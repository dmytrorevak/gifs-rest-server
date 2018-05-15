package main

import (
	"fmt"
	"github.com/dmytrorevak/gifs-rest-server/database"
	"github.com/dmytrorevak/gifs-rest-server/models"
)




func main() {
	db, err := database.InitDB("gifs_db.db")
	if err != nil {
		fmt.Printf("Error while init db: %v", err)
		return
	}
	//d.Migrate(db)
	gif := models.Gif{Age:11, }
	database.CreateGif(db, &gif)

	database.GetGif(db)
	//fmt.Print(a)
	// g := m.Gif{Age: 7}
	// d.CreateGif(db, &g)
	//d.GetGif(db)
}
