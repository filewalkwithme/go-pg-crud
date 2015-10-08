package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//Create
	var bookID int
	err = db.QueryRow(`INSERT INTO books(name, author, pages)
VALUES('Fight Club', 'Chuck Palahniuk', 208) RETURNING id`).Scan(&bookID)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\n", bookID)

	//Retrieve

	//Update

	//Delete

}
