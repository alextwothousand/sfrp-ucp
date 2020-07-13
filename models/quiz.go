package models

import "github.com/jinzhu/gorm"

// Quiz is the model for player's quiz questions that've been stored.
type Quiz struct {
	gorm.Model
	PlayerSQLID int    `gorm:"type:int;UNIQUE"`
	Questions   string `gorm:"type:json"`
	Answers     string `gorm:"type:json"`
}
