package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	d "gifs-rest-server/gifs-server/database"
	m "gifs-rest-server/gifs-server/models"
)




func main() {
	db, err := gorm.Open("mysql", "gifs_user:Gifs123456@/gifs_db")
	if err != nil {
		fmt.Print("Failed db connection")
	}
	defer db.Close()
	d.Migrate(db)
	d.GetGif(db, &m.Gif{})
	//fmt.Print(a)
	// g := m.Gif{Age: 7}
	// d.CreateGif(db, &g)
	//d.GetGif(db)
}
