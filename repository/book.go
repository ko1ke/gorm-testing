package repository

import (
	"github.com/moritamori/gorm-testing/model"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

type BookRepository interface {
	Create(book *model.Book) error
}

func (bookRepo BookRepositoryImpl) Create(book *model.Book) error {
	cx := bookRepo.DB.Create(book)
	return cx.Error
}
