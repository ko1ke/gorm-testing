package repository

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
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
		newId := 1
		rows := sqlmock.NewRows([]string{"id"}).AddRow(newId)
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(
			regexp.QuoteMeta(
				`INSERT INTO "books" ("created_at",` +
					`"updated_at","deleted_at","title",` +
					`"author") VALUES ($1,$2,$3,$4,$5) ` +
					`RETURNING "id"`),
		).WillReturnRows(rows)
		suite.mock.ExpectCommit()
		book := &model.Book{
			Title:  "Go言語の本",
			Author: "誰か",
		}
		err := suite.bookRepository.Create(book)

		if err != nil {
			suite.Fail("Error発生")
		}
		if book.ID != uint(newId) {
			suite.Fail("登録されるべきIDと異なっている")
		}
	})
}
