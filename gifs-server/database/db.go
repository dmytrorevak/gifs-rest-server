package database

import (
	"github.com/jinzhu/gorm"
	m "gifs-rest-server/gifs-server/models"
	"fmt"
)

func Migrate(db *gorm.DB)  {
	db.AutoMigrate(&m.Gif{})
}

func GetGif(db *gorm.DB) {
	a := db.First(&m.Gif{})
	fmt.Print(a.Rows())
}

func CreateGif(db *gorm.DB, gif *m.Gif)  {
	db.Create(&gif)
}
