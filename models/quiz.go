package models

import "github.com/jinzhu/gorm"

// Quiz is the model for player's quiz questions that've been stored.
type Quiz struct {
	gorm.Model
	UserID int
	RP     string `gorm:"type:text"`
	Rules  string `gorm:"type:text"`
}
