package models

type Book struct {
	Id     uint   `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}
