package main

import (
	"fmt"
	"library-api/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("GET /books", handlers.GetBooks)
	http.HandleFunc("GET /books/{id}", handlers.GetBookByID)
	http.HandleFunc("POST /books", handlers.CreateBook)

	fmt.Println("Server started on :8080")
	fmt.Println("Acess: http://localhost:8080/books")
	http.ListenAndServe(":8080", nil)

}
