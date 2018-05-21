package storage

import "github.com/jinzhu/gorm"

type Gif struct {
	gorm.Model
	Age int
}
