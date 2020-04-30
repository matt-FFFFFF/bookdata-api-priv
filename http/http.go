package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matt-FFFFFF/bookstore"
)

type api struct {
	Version string `json:"version"`
}

var apiInfo = api{
	Version: "1.0.0",
}

// Handler is the HTTP handler for the bookstore
type Handler struct {
	BookService bookstore.BookService
	AddressPort string
}

func (h *Handler) ServeHTTP() error {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", h.apiInfo).Methods(http.MethodGet)
	api.HandleFunc("/books", h.books).Methods(http.MethodGet)
	/* 	api.HandleFunc("/books/authors/{author}", searchByAuthor).Methods(http.MethodGet)
	   	api.HandleFunc("/books/book-name/{bookName}", searchByBookName).Methods(http.MethodGet)
	   	api.HandleFunc("/book/isbn/{isbn}", searchByISBN).Methods(http.MethodGet)
	   	api.HandleFunc("/book/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	   	api.HandleFunc("/book", createBook).Methods(http.MethodPost) */
	return http.ListenAndServe(h.AddressPort, r)
}

func (h *Handler) apiInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ai, err := json.Marshal(apiInfo)
	if err != nil {
		log.Println("could not marshal api info")
	}
	w.Write(ai)
	return
}

func (h *Handler) books(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//limit, err := getLimitParam(r)
	/* 	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
		return
	} */
	data, err := h.BookService.Books()
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}
