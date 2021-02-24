package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/moritamori/gorm-testing/model"
	"github.com/moritamori/gorm-testing/repository"
)

func main() {
	db, err := gorm.Open("postgres", "dbname=gormtesting password=mypassword")
	if err != nil {
		panic(err)
	}
	bookRepository := repository.BookRepositoryImpl{DB: db}

	book := model.Book{
		Title:  "Go言語の本",
		Author: "誰か",
	}
	newBook, err := bookRepository.Create(book)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Title: %s, Author: %s\n", newBook.Title, newBook.Author)
	}
}
