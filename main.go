package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadFile("./index.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(buf))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
