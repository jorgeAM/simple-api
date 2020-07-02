package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/jorgeAM/api/models"
	"github.com/stretchr/testify/assert"
)

const (
	sqlInsert          = "^INSERT INTO `users` (.+) VALUES"
	sqlSelectAll       = "^SELECT (.+) FROM `users`$"
	sqlSelectWithWhere = "^SELECT (.+) FROM `users` WHERE (.+) ORDER BY `users`.`id` ASC LIMIT 1$"
	sqlUpdate          = "^UPDATE `users` SET (.+) WHERE (.+)"
	sqlDelete          = "DELETE FROM `users` WHERE (.+)"
)

func TestNewUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	gDB, err := gorm.Open("mysql", db)
	defer gDB.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	uRepo := &UserRepository{
		DB: gDB,
	}

	user := &models.User{
		ID:        1,
		Username:  "jorgeAM",
		FirstName: "jorge",
		LastName:  "alfaro",
	}

	mock.ExpectBegin()
	mock.ExpectExec(sqlInsert).WithArgs(
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
	).WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	mock.ExpectCommit()

	u, err := uRepo.NewUser(user)

	assert.Nilf(t, err, "%v Should be nil", err)
	assert.Equal(t, user, u)
}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	gDB, err := gorm.Open("mysql", db)
	defer gDB.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	uRepo := &UserRepository{
		DB: gDB,
	}

	rows := sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow(1, "jorgeAM", "jorge", "alfaro").
		AddRow(2, "liliMA", "liliana", "murga")

	usersExpected := []*models.User{
		{
			ID:        1,
			Username:  "jorgeAM",
			FirstName: "jorge",
			LastName:  "alfaro",
		},
		{
			ID:        2,
			Username:  "liliMA",
			FirstName: "liliana",
			LastName:  "murga",
		},
	}

	mock.ExpectQuery(sqlSelectAll).WillReturnRows(rows)
	users, err := uRepo.GetUsers()

	assert.Nilf(t, err, "%v Should be nil", err)
	assert.Equal(t, usersExpected, users)
}

func TestGetUsersWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	gDB, err := gorm.Open("mysql", db)
	defer gDB.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	uRepo := &UserRepository{
		DB: gDB,
	}

	mock.ExpectQuery(sqlSelectAll).WillReturnError(errors.New("something got wrong"))
	users, err := uRepo.GetUsers()

	assert.NotNilf(t, err, "%v Should not be nil", err)
	assert.Nilf(t, users, "%v Should be nil", users)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	gDB, err := gorm.Open("mysql", db)
	defer gDB.Close()

	uRepo := &UserRepository{
		DB: gDB,
	}

	user := &models.User{
		ID:        1,
		Username:  "jorgeAM",
		FirstName: "jorge",
		LastName:  "alfaro",
	}

	rows := sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow(user.ID, user.Username, user.FirstName, user.LastName)

	mock.ExpectQuery(sqlSelectWithWhere).WillReturnRows(rows)
	u, err := uRepo.GetUser(1)

	assert.Nilf(t, err, "%v Should be nil", err)
	assert.Equal(t, user, u)
}

func TestGetUserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	gDB, err := gorm.Open("mysql", db)
	defer gDB.Close()

	uRepo := &UserRepository{
		DB: gDB,
	}

	rows := sqlmock.NewRows(nil)

	mock.ExpectQuery(sqlSelectWithWhere).WillReturnRows(rows)
	u, err := uRepo.GetUser(1)

	assert.NotNil(t, err, "%v Should not be nil", err)
	assert.Nil(t, u, "%v should be nil", u)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	gDB, err := gorm.Open("mysql", db)
	defer gDB.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	uRepo := &UserRepository{
		DB: gDB,
	}

	user := &models.User{
		ID:        1,
		Username:  "jorgeAM",
		FirstName: "jorge",
		LastName:  "alfaro",
	}

	sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow(user.ID, user.Username, user.FirstName, user.LastName)

	mock.ExpectBegin()
	mock.ExpectExec(sqlDelete).WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	mock.ExpectCommit()

	err = uRepo.DeleteUser(user.ID)

	assert.Nilf(t, err, "%v Should be nil", err)
}

func TestDeleteUserWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	gDB, err := gorm.Open("mysql", db)
	defer gDB.Close()

	assert.Nilf(t, err, "%v Should be nil", err)

	uRepo := &UserRepository{
		DB: gDB,
	}

	mock.ExpectBegin()
	mock.ExpectExec(sqlDelete).WillReturnResult(sqlmock.NewErrorResult(errors.New("not found")))
	mock.ExpectCommit()

	err = uRepo.DeleteUser(1)

	assert.NotNilf(t, err, "%v Should not be nil", err)
}
