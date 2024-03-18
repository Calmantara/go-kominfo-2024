package repository

import (
	"context"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/infrastructure/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMockGorm() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}
	return gormDB, mock
}
func TestGetUsers(t *testing.T) {
	t.Run("error get users", func(t *testing.T) {
		db, mock := newMockGorm()
		// mock infra
		postgresMock := mocks.NewGormPostgres(t)
		postgresMock.On("GetConnection").Return(db)
		// mock query
		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT * FROM "users"
		`)).WillReturnError(errors.New("some error"))

		userRepo := userQueryImpl{db: postgresMock}
		res, err := userRepo.GetUsers(context.Background())
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(res))
	})

	t.Run("success get users", func(t *testing.T) {
		db, mock := newMockGorm()
		// mock infra
		postgresMock := mocks.NewGormPostgres(t)
		postgresMock.On("GetConnection").Return(db)
		// mock query
		row := sqlmock.
			NewRows([]string{"id", "username"}).
			AddRow(1, "username")

		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT * FROM "users"
		`)).WillReturnRows(row)

		userRepo := userQueryImpl{db: postgresMock}
		res, err := userRepo.GetUsers(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
	})
}
