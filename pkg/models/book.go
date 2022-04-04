package models

import (
	//import gorm
	"github.com/jinzhu/gorm"

	"github.com/lucho-sosa/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name  string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication   string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.getDB()
	db.AutoMigrate(&Book{})
}




