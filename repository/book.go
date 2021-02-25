package repository

import (
	"github.com/moritamori/gorm-testing/model"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

type BookRepository interface {
	Create(string, string) (*model.Book, error)
}

func (bookRepo BookRepositoryImpl) Create(book model.Book) (*model.Book, error) {
	err := bookRepo.DB.Create(&book).Error
	return &book, err
}
