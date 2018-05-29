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
		return nil, err
	}

	// Creating the table
	if !db.HasTable(&Gif{}) {
		if err = db.CreateTable(&Gif{}).Set("gorm:table_options", "ENGINE=InnoDB").Error; err != nil {
			fmt.Printf("Error creation Gifs table: %v", err)
			return nil, err
		}
	}

	return db, nil
}

func GetGifsList(db *gorm.DB) ([]Gif, error) {
	var gifs GifsList
	return gifs, db.Find(&gifs).Error
}

func GetGif(db *gorm.DB, id int) (Gif, error) {
	var g Gif
	return g, db.First(&g, id).Error
}

func CreateGif(db *gorm.DB, gif *Gif) error {
	return db.Create(&gif).Error
}

func UpdateGif(db *gorm.DB, id int, data *Gif) error {
	var g Gif
	err := db.First(&g, id).Error
	if err != nil {
		return err
	}
	// TODO: validate input data before update
	err = db.Model(&g).Updates(data).Error
	if err != nil {
		// handle error higher
		return err
	}
	return db.Save(&g).Error
}

func DeleteGif(db *gorm.DB, id int) error {
	var g Gif
	err := db.First(&g, id).Error
	if err != nil {
		return err
	}
	return db.Unscoped().Delete(&g).Error
}
