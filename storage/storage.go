package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sanfierrorp/ucp/config"
	"github.com/sanfierrorp/ucp/models"

	// Adds in the MySQL/MariaDB dialect.
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var stDb *gorm.DB = nil

// Instance returns a instance of the gorm database.
func Instance() *gorm.DB {
	return stDb
}

// Open allows you to open up a connection with the database.
func Open() {
	var stCfg = config.Database

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", stCfg.Username, stCfg.Password, stCfg.Database))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.News{})
	db.AutoMigrate(&models.Players{})
	db.AutoMigrate(&models.Quiz{})

	stDb = db
	fmt.Println("connected")
}

// Close allows you to close the connection with the database.
func Close() error {
	if stDb == nil {
		return fmt.Errorf("Database handle not online")
	}

	err := stDb.Close()
	if err != nil {
		return err
	}

	fmt.Println("\ndisconnecting")
	return nil
}
