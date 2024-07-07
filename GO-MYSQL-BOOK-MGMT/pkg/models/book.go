package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rajjmalla10/GO-PROJECT-UPDATED/GO-MYSQL-BOOK-MGMT/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if db == nil {
		log.Fatal("Failed to initialize database connection")
	}

	if err := db.AutoMigrate(&Book{}).Error; err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func (b *Book) CreateBook() *Book {
	if db.NewRecord(b) { // Checks if 'b' is a new record in the database
		db.Create(&b) // Creates a new record in the 'books' table with data from 'b'
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
