package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/moritamori/gorm-testing/model"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

type BookRepository interface {
	Create(string, string) (*model.Book, error)
}

func (bookRepository BookRepositoryImpl) Create(title string, author string) (*model.Book, error) {
	book := model.Book{
		Title:  title,
		Author: author,
	}
	err := bookRepository.DB.Create(&book).Error
	return &book, err
}
