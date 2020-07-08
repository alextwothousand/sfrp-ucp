package models

import "github.com/jinzhu/gorm"

// Users is the model for users registered on the UCP.
type Users struct {
	gorm.Model
	Username string
	Password string
	Email    string
}
