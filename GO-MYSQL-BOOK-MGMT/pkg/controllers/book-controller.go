package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rajjmalla10/GO-PROJECT-UPDATED/GO-MYSQL-BOOK-MGMT/pkg/models"
	"github.com/rajjmalla10/GO-PROJECT-UPDATED/GO-MYSQL-BOOK-MGMT/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing the Id")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}

	// Parse request body into the newBook instance using utils.ParseBody
	if err := utils.ParseBody(r, CreateBook); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Create the book in the database using the CreateBook method
	b := CreateBook.CreateBook()

	// Marshal the createdBook into JSON
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON-encoded data to the response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing the bookId")
	}
	del := models.DeleteBook(ID)
	res, _ := json.Marshal(del)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing", err)
	}
	get, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		get.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		get.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		get.Author = updateBook.Publication

	}
	db.Save(&get)
	res, _ := json.Marshal(get)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
