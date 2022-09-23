package postgres

import (
	"KatodForAnod/siteGallery/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

var (
	testPostgres postgreSQl
)

func TestPostgreSQl_AddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUser := models.User{
		User:     "testname",
		Email:    "test@mail.ru",
		PassHash: "",
	}

	testPostgres.conn = db
	mock.ExpectExec("INSERT INTO users").WithArgs(testUser.User, testUser.Email, testUser.PassHash)

	if err = testPostgres.AddUser(testUser); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
