package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB(sqlStoragePath string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", sqlStoragePath)
	if err != nil {
		fmt.Print(err)
	}

	// Creating the table
	if !db.HasTable(&Gif{}) {
		if err = db.CreateTable(&Gif{}).Set("gorm:table_options", "ENGINE=InnoDB").Error; err != nil {
			fmt.Print(err)
		}
	}

	return db, nil
}

func GetGif(db *gorm.DB) Gif {
	var g Gif
	err := db.Find(&g).Error
	if err != nil {
		fmt.Printf("Error getting gif:%v", err)
	}

	return g
}

func GetGifsList(db *gorm.DB) []Gif {
	var g []Gif
	err := db.Find(&g).Error
	if err != nil {
		fmt.Printf("Error getting gif:%v", err)
	}

	return g
}

func CreateGif(db *gorm.DB, gif *Gif) {
	err := db.Create(&gif).Error
	if err != nil {
		fmt.Printf("Error creating gif:%v", err)
	}
}
