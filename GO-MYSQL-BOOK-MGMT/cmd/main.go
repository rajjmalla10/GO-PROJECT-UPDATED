package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rajjmalla10/GO-PROJECT-UPDATED/GO-MYSQL-BOOK-MGMT/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	port := 9010
	e := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if e != nil {
		fmt.Println("error starting server...", e)
		return
	}

}
