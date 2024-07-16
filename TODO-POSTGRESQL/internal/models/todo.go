package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rajjmalla10/TODO-POSTGRESQL/internal/config"
)

type Todo struct {
	Title string `json:"title"`
}

var db *sql.DB

func init() {

	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer db.Close()
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT title, completed FROM todos")
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Println("Failed to Connect to database", err)
		return
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Title)
		if err != nil {
			log.Println("error scanning todo", err)
			continue
		}
		todos = append(todos, todo)

	}

	//Searlizing data to json

	jsonData, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "Failed to searlizie todos", http.StatusInternalServerError)
		log.Println("failed to serialize todos:", err)
		return
	}

	w.Header().Set("content-type", "applcation/json")
	w.Write(jsonData)
}

func PostTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Raj Malla")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
