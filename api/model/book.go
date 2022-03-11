package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `validate:"required"`
	Author string `validate:"required"`
}
