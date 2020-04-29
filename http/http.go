package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matt-FFFFFF/bookstore"
)

type Handler struct {
	BookService bookstore.BookService
}

func (h *Handler) ServeHTTP() error {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", h.ApiInfo).Methods(http.MethodGet)
	api.HandleFunc("/books", h.Books).Methods(http.MethodGet)
	/* 	api.HandleFunc("/books/authors/{author}", searchByAuthor).Methods(http.MethodGet)
	   	api.HandleFunc("/books/book-name/{bookName}", searchByBookName).Methods(http.MethodGet)
	   	api.HandleFunc("/book/isbn/{isbn}", searchByISBN).Methods(http.MethodGet)
	   	api.HandleFunc("/book/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	   	api.HandleFunc("/book", createBook).Methods(http.MethodPost) */
	return http.ListenAndServe(":8080", r)
}

func (h *Handler) ApiInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"api_version": "1.0.0"}`))
	return
}

func (h *Handler) Books(w http.ResponseWriter, r *http.Request) {
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

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}

		limit = val
	}
	return limit, nil
}
