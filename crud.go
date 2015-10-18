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

	//Update
	var bookID int
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

func insertBook(name, author string, pages int, publicationDate time.Time) int {
	db, err := sql.Open("postgres", "user=postgres dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//Create
	var bookID int
	err = db.QueryRow(`INSERT INTO books(name, author, pages, publication_date)
VALUES($1, $2, $3, $4) RETURNING id`, name, author, pages, publicationDate).Scan(&bookID)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Last inserted ID: %v\n", bookID)
	return bookID
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

func getBook(bookID int) book {
	db, err := sql.Open("postgres", "user=postgres dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//Retrieve
	res := book{}

	var id int
	var name string
	var author string
	var pages int
	var publicationDate pq.NullTime

	err = db.QueryRow(`SELECT id, name, author, pages, publication_date FROM books where id = $1`, bookID).Scan(&id, &name, &author, &pages, &publicationDate)
	if err == nil {
		res = book{ID: id, Name: name, Author: author, Pages: pages, PublicationDate: publicationDate.Time}
	}
	return res
}
