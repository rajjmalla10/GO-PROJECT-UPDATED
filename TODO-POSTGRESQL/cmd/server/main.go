package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/rajjmalla10/TODO-POSTGRESQL/internal/routes"
)

func main() {

	r := mux.NewRouter()
	routes.TodoRoutes(r)
	http.Handle("/", r)

	engine := template.Must(template.ParseGlob("./internal/views/*.html"))

	port := 3000
	e := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if e != nil {
		fmt.Println("error starting server...", e)
		return
	}

}
