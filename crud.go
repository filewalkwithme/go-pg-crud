package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

func crud() {
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
	fmt.Printf("Last inserted ID: %v\n", bookID)

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
				log.Fatalf("err: %v\n", err)
			}

		}
	} else {
		log.Fatalf("err: %v\n", err)
	}

	//Update
	var newPublicationDate = time.Date(1996, time.August, 17, 0, 0, 0, 0, time.UTC)
	res, err := db.Exec(`UPDATE books SET publication_date = $1 where id = $2`, newPublicationDate, bookID)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	fmt.Printf("Number of rows updated: %v\n", rowsUpdated)

	//Delete
	res, err = db.Exec(`delete from books where id = $1`, bookID)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	fmt.Printf("Number of rows deleted: %v\n", rowsDeleted)
}

func allBooks() []book {
	db, err := sql.Open("postgres", "user=postgres dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//Retrieve
	books := []book{}

	rows, err := db.Query(`SELECT id, name, author, pages, publication_date FROM books order by id`)
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
				currentBook := book{ID: id, Name: name, Author: author, Pages: pages}
				if publicationDate.Valid {
					currentBook.PublicationDate = publicationDate.Time
				}

				books = append(books, currentBook)
			} else {
				log.Fatalf("err: %v\n", err)
			}

		}
	} else {
		log.Fatalf("err: %v\n", err)
	}

	return books
}
