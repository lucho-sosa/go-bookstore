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
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db:= db.Where("ID = ?", id).First(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(&Book{})
	return book
}


