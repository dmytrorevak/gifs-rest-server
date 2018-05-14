package models

import "github.com/jinzhu/gorm"

type Gif struct {
	gorm.Model
	Title string
}
