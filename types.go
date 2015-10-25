package main

import "time"

//IndexPage represents the content of the index page, available on "/"
//The index page shows a list of all books stored on db
type IndexPage struct {
	AllBooks []Book
}

//BookPage represents the content of the book page, available on "/book.html"
//The book page shows info about a given book
type BookPage struct {
	TargetBook Book
}

//Book represents a book object
type Book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}

//PublicationDateStr returns a sanitized Publication Date in the format YYYY-MM-DD
func (b Book) PublicationDateStr() string {
	return b.PublicationDate.Format("2006-01-02")
}

//ErrorPage represents shows an error message, available on "/book.html"
type ErrorPage struct {
	ErrorMsg string
}
