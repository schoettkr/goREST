package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock id - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			// Create updated/new book from request body
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock data - @todo - Implement in DB
	books = append(books, Book{ID: "1", Isbn: "348954", Title: "Bible", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	book2 := Book{ID: "2", Isbn: "90834", Title: "Book Two", Author: &Author{Firstname: "Max", Lastname: "Mustermann"}}
	books = append(books, book2)

	// Router Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

	testTR := triangle{2, 6, 3}
	println("Perimeter is ", testTR.perimeter())
}

type triangle struct {
	a int
	b int
	c int
}

func (tr triangle) perimeter() int {
	return tr.a + tr.b + tr.c
}
