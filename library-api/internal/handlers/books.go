package handlers

import (
	"encoding/json"
	m "library-api/internal/models"
	"library-api/internal/store"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store.Books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido: deve ser numérico", http.StatusBadRequest)
		return
	}

	for _, book := range store.Books {
		if int(book.Id) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Livro não encontrado", http.StatusNotFound)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book m.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	for _, b := range store.Books {
		if b.Title == book.Title || b.Author == book.Author {
			http.Error(w, "O item já existe!", http.StatusBadRequest)
			return
		}
	}

	newID := len(store.Books) + 1
	book.Id = uint(newID)

	store.Books = append(store.Books, book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
