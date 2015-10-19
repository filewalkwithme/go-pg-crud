package main

import "time"

type indexPage struct {
	Books []book
}

type bookPage struct {
	Book book
}

type book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}
