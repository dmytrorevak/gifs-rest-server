package storage

import (
    "github.com/jinzhu/gorm"
    "fmt"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&Gif{})
}

func InitDB(sqlStoragePath string) (*gorm.DB, error) {
    db, err := gorm.Open("mysql", sqlStoragePath)
    if err != nil {
        fmt.Print(err)
    }

    // Display SQL queries
    // if err = db.LogMode(true).Error; err != nil {
        // return nil, err
    // }

    // Creating the table
    if !db.HasTable(&Gif{}) {
        if err = db.CreateTable(&Gif{}).Set("gorm:table_options", "ENGINE=InnoDB").Error; err != nil {
            fmt.Print(err)
        }
        // Migrate(db)
    }

    // db = db.Debug()

    return db, nil
}

func GetGif(db *gorm.DB) Gif {
    var g Gif
    err := db.Find(&g).Error
    if err != nil {
        fmt.Printf("Error getting gif:%v", err)
    }
    // fmt.Printf("GIFS:%v", g)
    // var gif Gif
    // err = db.First(&g).Error
    // if err != nil {
    //     fmt.Printf("Error getting gif:%v", err)
    // }
    // fmt.Printf("First GIF:%v", gif)
    return g
}

func GetGifsList(db *gorm.DB) []Gif {
    var g []Gif
    err := db.Find(&g).Error
    if err != nil {
        fmt.Printf("Error getting gif:%v", err)
    }
    // fmt.Printf("GIFS:%v", g)
    // var gif Gif
    // err = db.First(&g).Error
    // if err != nil {
    //     fmt.Printf("Error getting gif:%v", err)
    // }
    // fmt.Printf("First GIF:%v", gif)
    return g
}

func CreateGif(db *gorm.DB, gif *Gif) {
    // fmt.Println("create gif\n")
    err := db.Create(&gif).Error
    if err != nil {
                fmt.Printf("Error creating gif:%v", err)
        }
}
