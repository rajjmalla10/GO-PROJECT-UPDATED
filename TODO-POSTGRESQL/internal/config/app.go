package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	connstr := "postgresql://postgres:password@localhost/todos?sslmode=disable"

	//connect to the database

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connected to postgreSQL database!")
	return db, nil
}
