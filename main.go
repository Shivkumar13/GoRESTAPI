package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	isbn   string  `json:"isbn"`
	Title  string  `json:"Title"`
	Author *Author `json:"Author"`
}

//Author Struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
}

//Init books var as a `slice` Book struct

var Books []Book

//Get Single book

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params

	// Loop through books and find with id
	for _, item := range Books {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return

		}
	}

	json.NewEncoder(w).Encode(&Book{})

}

//Get all Books
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)

}

// create Book

func createBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Book

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

// Update Book

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//Init Router
	r := mux.NewRouter()

	//Mock Data - @todo - implement a DB

	Books = append(Books, Book{ID: "1", isbn: "130797", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	Books = append(Books, Book{ID: "2", isbn: "132553", Title: "Book Two", Author: &Author{Firstname: "Shivkumar", Lastname: "Ople"}})
	Books = append(Books, Book{ID: "3", isbn: "342355", Title: "Book Three", Author: &Author{Firstname: "Sherlock", Lastname: "Holmes"}})

	//Route Handlers / Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
