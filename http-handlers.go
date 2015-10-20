package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func handleListBooks(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("www/index.html")
	if err != nil {
		log.Fatal(err)
	}

	var page = IndexPage{AllBooks: allBooks()}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}

func handleViewBook(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("www/book.html")
	if err != nil {
		log.Fatal(err)
	}

	v := r.URL.Query()
	idStr := v.Get("id")
	currentBook := Book{}
	if len(idStr) > 0 {
		id, _ := strconv.Atoi(idStr)

		currentBook = getBook(id)
	}

	var page = BookPage{TargetBook: currentBook}
	bookPage := string(buf)
	t := template.Must(template.New("bookPage").Parse(bookPage))
	t.Execute(w, page)
}

func handleSaveBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	v := r.PostForm
	idStr := v.Get("id")
	id := 0
	if len(idStr) > 0 {
		id, _ = strconv.Atoi(idStr)
	}

	name := v.Get("name")
	author := v.Get("author")

	pagesStr := v.Get("pages")
	pages := 0
	if len(pagesStr) > 0 {
		pages, _ = strconv.Atoi(pagesStr)
	}

	publicationDateStr := v.Get("publicationDate")
	var publicationDate time.Time

	if len(publicationDateStr) > 0 {
		publicationDate, _ = time.Parse("2006-01-02", publicationDateStr)
	}

	if id == 0 {
		insertBook(name, author, pages, publicationDate)
	} else {
		updateBook(id, name, author, pages, publicationDate)
	}

	http.Redirect(w, r, "/", 302)
}

func handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			n, err := removeBook(id)
			if err == nil {
				fmt.Printf("Rows removed: %v\n", n)
				http.Redirect(w, r, "/", 302)
			} else {
				renderErrorPage(w, err)
				return
			}
		} else {
			renderErrorPage(w, err)
			return
		}
	}
}

func renderErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("www/error.html")
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
