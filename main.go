package main

import (
	"database/sql"
	"fmt"
	"log"

	pq "github.com/lib/pq"
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
	rows, err := db.Query(`SELECT id, name, author, pages, publication_date FROM books WHERE id = $1`, bookID)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var author string
			var pages int
			var publicationDate pq.NullTime

			err = rows.Scan(&id, &name, &author, &pages, &publicationDate)
			if err == nil {
				fmt.Printf("id: %v\n", id)
				fmt.Printf("name: %v\n", name)
				fmt.Printf("author: %v\n", author)
				fmt.Printf("pages: %v\n", pages)
				if publicationDate.Valid {
					fmt.Printf("publicationDate: %v\n", publicationDate.Time)
				} else {
					fmt.Printf("publicationDate: null\n")
				}
			} else {
				fmt.Printf("err: %v\n", err)
			}

		}
	} else {
		fmt.Printf("err: %v\n", err)
	}

	//Update

	//Delete

}
