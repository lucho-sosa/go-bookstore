package config

import (
	// import gorm 
 "github.com/jinzhu/gorm/"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:lucho90@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}