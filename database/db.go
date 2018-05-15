package database

import (
	"github.com/jinzhu/gorm"
	"github.com/dmytrorevak/gifs-rest-server/models"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	// "encoding/json"
)

func Migrate(db *gorm.DB)  {
	db.AutoMigrate(&models.Gif{})
}

func InitDB(sqlStoragePath string) (*gorm.DB, error) {
	// Opening file
	db, err := gorm.Open("sqlite3", sqlStoragePath)
	if err != nil {
		return nil, err
	}
	// Display SQL queries
	if err = db.LogMode(true).Error; err != nil {
		return nil, err
	}
	// Creating the table
	if !db.HasTable(&models.Gif{}) {
		if err = db.CreateTable(&models.Gif{}).Set("gorm:table_options", "ENGINE=InnoDB").Error; err != nil {
			return nil, err
		}
	}

	db = db.Debug()

	return db, nil
}

func GetGif(db *gorm.DB) {
	var g []models.Gif
	err := db.Find(&g).Error
	if err != nil {
		fmt.Printf("Error getting gif:%v", err)
	}
	fmt.Printf("GIFS:%v", g)
	var gif models.Gif
	err = db.First(&g).Error
	if err != nil {
		fmt.Printf("Error getting gif:%v", err)
	}
	fmt.Printf("First GIF:%v", gif)
}

func CreateGif(db *gorm.DB, gif *models.Gif)  {
	fmt.Println("create gif\n")
	err := db.Create(&gif).Error
	if err != nil {
                fmt.Printf("Error creating gif:%v", err)
        }
}
