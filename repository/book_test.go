package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moritamori/gorm-testing/model"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// テストスイートの構造体
type BookRepositoryTestSuite struct {
	suite.Suite
	bookRepository BookRepositoryImpl
	mock           sqlmock.Sqlmock
}

// テストのセットアップ
// (sqlmockをNew、Gormで発行されるクエリがモックに送られるように)
func (suite *BookRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	bookRepository := BookRepositoryImpl{}
	bookRepository.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	suite.bookRepository = bookRepository
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *BookRepositoryTestSuite) TearDownTest() {
	db, _ := suite.bookRepository.DB.DB()
	db.Close()
}

// テストスイートの実行
func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
}

// Createのテスト
func (suite *BookRepositoryTestSuite) TestCreate() {
	suite.Run("create a book", func() {
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
