package controllers

import "learning_gin/project/models"

var Library []models.Book

var Counter int

func InitDatabase() {
	Counter = 1

	book1 := models.Book{
		Id:     1,
		Title:  "Généalogie de la morale",
		Author: "Friedrich Nietzsche",
	}
	Library = append(Library, book1)
}
