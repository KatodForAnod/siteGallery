package postgres

import (
	"KatodForAnod/siteGallery/internal/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
	mock.ExpectExec("INSERT INTO users").WithArgs(testUser.Email, testUser.User, testUser.PassHash).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err = testPostgres.AddUser(testUser); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgreSQl_AddImage(t *testing.T) {

}

func TestPostgreSQl_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUser := models.User{
		Id:       0,
		User:     "testName",
		Email:    "test@mail.ru",
		PassHash: "",
	}

	testPostgres.conn = db
	rows := sqlmock.NewRows([]string{"email", "id", "name", "password"}).
		AddRow(testUser.Email, testUser.Id, testUser.User, "")
	mock.ExpectQuery("SELECT email, id, name, password FROM users WHERE email = ").
		WithArgs(testUser.Email).WillReturnRows(rows)

	userResp, err := testPostgres.GetUser(testUser.Email)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	assert.Equal(t, testUser, userResp, "returned struct not correct")
}

func TestPostgreSQl_GetUserErr(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testPostgres.conn = db
	mock.ExpectQuery("SELECT email, id, name, password FROM users WHERE email = ").
		WithArgs("").WillReturnError(fmt.Errorf("user not found"))

	if _, err = testPostgres.GetUser(""); err == nil {
		t.Errorf("error was expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
