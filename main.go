package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "dbname=books_database user=postgres password=postgres host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListBooks)
	http.HandleFunc("/book.html", handleViewBook)
	http.HandleFunc("/save", handleSaveBook)
	http.HandleFunc("/delete", handleDeleteBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
