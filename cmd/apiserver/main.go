package main

import (
	"log"

	"github.com/matt-FFFFFF/bookstore/datastore/memory"
	"github.com/matt-FFFFFF/bookstore/http"
)

func main() {
	s, err := memory.Init("filepath")
	if err != nil {
		log.Fatalln("could not init memory-backed store:", err)
	}

	var h http.Handler
	h.BookService = s
	log.Fatalln(h.ServeHTTP())
}
