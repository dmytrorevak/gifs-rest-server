package database

import (
	"github.com/jinzhu/gorm"
	m "gifs-rest-server/gifs-server/models"
	"fmt"
	// "encoding/json"
)

func Migrate(db *gorm.DB)  {
	db.AutoMigrate(&m.Gif{})
}

func GetGif(db *gorm.DB, gif *m.Gif) {
	var g m.Gif
	db.First(&g)
	// json.Unmarshal(&g, m.Gif)
	fmt.Print(g)
}

func CreateGif(db *gorm.DB, gif *m.Gif)  {
	db.Create(&gif)
}
