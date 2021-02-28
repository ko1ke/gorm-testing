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
	book := &model.Book{
		Title:  "Go言語の本",
		Author: "誰か",
	}
	err = bookRepo.Create(book)

	// エラーが発生しないかチェック
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success!")
}
