package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/moritamori/gorm-testing/model"
	"github.com/stretchr/testify/suite"
)

type BookRepositoryTestSuite struct {
	suite.Suite
	bookRepository BookRepositoryImpl
	mock           sqlmock.Sqlmock
}

func (suite *BookRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	bookRepository := BookRepositoryImpl{}
	bookRepository.DB, _ = gorm.Open("postgres", db)
	suite.bookRepository = bookRepository
}

func (suite *BookRepositoryTestSuite) TearDownTest() {
	suite.bookRepository.DB.Close()
}

func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
}

func (suite *BookRepositoryTestSuite) TestCreate() {
	suite.Run("書籍を登録", func() {
		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(`INSERT INTO "books"`).WillReturnRows(rows)
		suite.mock.ExpectCommit()
		book := model.Book{
			Title:  "Go言語の本",
			Author: "誰か",
		}
		actual, err := suite.bookRepository.Create(book)

		if err != nil {
			suite.Fail("Error発生")
		}
		if actual.Title != "Go言語の本" {
			suite.Fail("登録された書籍とタイトルが同じではない")
		}
		if actual.Author != "誰か" {
			suite.Fail("登録された書籍と著者が同じではない")
		}
	})
}
