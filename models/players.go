package models

import "github.com/jinzhu/gorm"

// Players is the model for users registered on the UCP.
type Players struct {
	gorm.Model
	Name         string `gorm:"UNIQUE"`
	Password     string
	Email        string
	RegisterIP   string
	IP           string
	QuizComplete bool `gorm:"type:boolean;default:'0';"`
	AdminLevel   int  `gorm:"default:'0'"`
	Donator      int  `gorm:"default:'0'"`
	Helper       int  `gorm:"default:'0'"`
}
