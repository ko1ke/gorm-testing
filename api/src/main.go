package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/moritamori/gorm-testing/model"
	"github.com/moritamori/gorm-testing/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	// DB接続を開く
	if err != nil {
		log.Printf("failed to load .env: %v", err)
	}
	dsnString := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsnString), &gorm.Config{})
	if err != nil {

		panic("failed to connect database:" + dsnString)
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
