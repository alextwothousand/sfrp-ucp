package models

import "github.com/jinzhu/gorm"

// Players is the model for users registered on the UCP.
type Players struct {
	gorm.Model
	Name       string
	Password   string
	Email      string
	RegisterIP string
	IP         string
	AdminLevel int
	Donator    int
	Helper     int
}
