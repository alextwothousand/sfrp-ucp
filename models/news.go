package models

import "github.com/jinzhu/gorm"

// News is the model for front-page news articles.
type News struct {
	gorm.Model
	Author string
	Text   string `gorm:"type:text"`
}
