package models

type Book struct {
	Id     int    `json:"id"`
	Title  string `string:"title"`
	Author string `string:"author"`
}
