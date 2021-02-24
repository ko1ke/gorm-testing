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

func (bookRepository BookRepositoryImpl) Create(book model.Book) (*model.Book, error) {
	err := bookRepository.DB.Create(&book).Error
	return &book, err
}
