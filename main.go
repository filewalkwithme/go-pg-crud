package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

type indexPage struct {
	Books []book
}

type book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadFile("www/index.html")
		if err != nil {
			log.Fatal(err)
		}

		var page = indexPage{Books: allBooks()}

		indexPage := string(buf)

		t := template.Must(template.New("indexPage").Parse(indexPage))

		t.Execute(w, page)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
