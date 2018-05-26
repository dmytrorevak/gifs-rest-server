package storage

import "github.com/jinzhu/gorm"

type Gif struct {
	gorm.Model
	Url     string
	Rating  int
	Name    string
	Size    float32
}

type GifsList []Gif
