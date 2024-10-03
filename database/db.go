package database

import (
	
	"github.com/ajay-kumar-ks/gin-test/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Initialise(){
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	DB.AutoMigrate(&models.User{})
}