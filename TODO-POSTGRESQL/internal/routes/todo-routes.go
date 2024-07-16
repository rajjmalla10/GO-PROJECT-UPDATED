package routes

import (
	"github.com/gorilla/mux"
	"github.com/rajjmalla10/TODO-POSTGRESQL/internal/models"
)

var TodoRoutes = func(router *mux.Router) {
	router.HandleFunc("/", models.GetTodo).Methods("GET")
	router.HandleFunc("/", models.PostTodo).Methods("POST")
	router.HandleFunc("/update", models.UpdateTodo).Methods("PUT")
	router.HandleFunc("/delete", models.DeleteTodo).Methods("DELETE")
}
