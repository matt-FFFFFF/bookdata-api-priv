package main

import (
	"log"

	"github.com/matt-FFFFFF/bookstore/http"
)

func main() {
	var h http.Handler
	log.Fatalln(h.ServeHTTP())
}
