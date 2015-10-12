package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadFile("www/index.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(buf))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
