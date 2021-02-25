package main

import (
	"fmt"

	"github.com/moritamori/gorm-testing/model"
	"github.com/moritamori/gorm-testing/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB接続を開く
	url := "dbname=gormtesting password=mypassword"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// リポジトリ(`repository/book.go`)を介して、書籍(book)のデータを登録
	bookRepo := repository.BookRepositoryImpl{DB: db}
	book := model.Book{
		Title:  "Go言語の本",
		Author: "誰か",
	}
	newBook, err := bookRepo.Create(book)

	// エラーが発生しなければ、登録した書籍データの内容を標準出力
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Title: %s, Author: %s\n", newBook.Title, newBook.Author)
	}
}
