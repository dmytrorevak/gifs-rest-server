package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB(sqlStoragePath string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", sqlStoragePath)
	if err != nil {
		fmt.Printf("Error init creation db: %v", err)
	}

	// Creating the table
	if !db.HasTable(&Gif{}) {
		if err = db.CreateTable(&Gif{}).Set("gorm:table_options", "ENGINE=InnoDB").Error; err != nil {
			fmt.Printf("Error creation Gifs table: %v", err)
		}
	}

	return db, nil
}

func GetGifsList(db *gorm.DB) []Gif {
	var gifs GifsList
	err := db.Find(&gifs).Error
	if err != nil {
		fmt.Printf("Error getting gifs list: %v", err)
	}

	return gifs
}

func GetGif(db *gorm.DB, id int) Gif {
	var g Gif
	err := db.First(&g, id).Error
	if err != nil {
		fmt.Printf("Error getting gif:%v", err)
	}

	return g
}

func CreateGif(db *gorm.DB, gif *Gif) {
	err := db.Create(&gif).Error
	if err != nil {
		fmt.Printf("Error creating gif: %v", err)
	}
}

func UpdateGif(db *gorm.DB, id int, data *Gif)  {
	var g Gif
	err := db.First(&g, id).Error
	if err != nil {
		fmt.Printf("Error getting gif: %v", err)
	}

	db.Model(&g).Updates(data)
	db.Save(&g)
}

func DeleteGif(db *gorm.DB, id int)  {
	var g Gif
	err := db.First(&g, id).Error
	if err != nil {
		fmt.Printf("Error getting gif:%v", err)
	}

	db.Unscoped().Delete(&g)
}
