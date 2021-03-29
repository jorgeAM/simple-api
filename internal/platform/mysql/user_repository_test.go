package mysql

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/jorgeAM/simple-api/internal/user/domain"
)

const (
	sqlInsert          = "^INSERT INTO `users` (.+) VALUES"
	sqlSelectAll       = "^SELECT (.+) FROM `users`$"
	sqlSelectWithWhere = "^SELECT (.+) FROM `users` WHERE (.+) ORDER BY `users`.`id` ASC LIMIT 1$"
	sqlDelete          = "DELETE FROM `users` WHERE (.+)"
)

type UserSuite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository domain.Repository
	user       *domain.User
}

func (s *UserSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(false)

	s.repository = NewUserRepository(s.DB)

	user, err := domain.NewUser("88109b71-996c-42cd-997e-cbf81cf8f886", "jorgeAM", "jorge", "alfaro")
	require.NoError(s.T(), err)

	s.user = user
}

func (s *UserSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserSuite) TestNewUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlInsert).WithArgs(
		s.user.ID.String(),
		s.user.Username.String(),
		s.user.FirstName.String(),
		s.user.LastName.String(),
	).WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	err := s.repository.NewUser(context.Background(), s.user)

	assert.Nilf(s.T(), err, "%v Should be nil", err)
}

func (s *UserSuite) TestNewUserWithError() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlInsert).WithArgs(
		s.user.ID.String(),
		s.user.Username.String(),
		s.user.FirstName.String(),
		s.user.LastName.String(),
	).WillReturnError(errors.New("Something got wrong to save record"))
	s.mock.ExpectRollback()

	err := s.repository.NewUser(context.Background(), s.user)

	assert.NotNilf(s.T(), err, "%v should not be nil", err)
}

func (s *UserSuite) TestGetUsers() {
	rows := sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow("88109b71-996c-42cd-997e-cbf81cf8f885", "jorgeAM", "jorge", "alfaro").
		AddRow("88109b71-996c-42cd-997e-cbf81cf8f881", "liliMA", "liliana", "murga")

	s.mock.ExpectQuery(sqlSelectAll).WillReturnRows(rows)

	_, err := s.repository.GetUsers(context.Background())

	assert.Nilf(s.T(), err, "%v Should be nil", err)
}

func (s *UserSuite) TestGetUsersWithError() {
	s.mock.ExpectQuery(sqlSelectAll).WillReturnError(errors.New("something got wrong"))

	users, err := s.repository.GetUsers(context.Background())

	assert.NotNilf(s.T(), err, "%v Should not be nil", err)
	assert.Nilf(s.T(), users, "%v Should be nil", users)
}

func (s *UserSuite) GetUser() {
	rows := sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow(s.user.ID.String(), s.user.Username.String(), s.user.FirstName.String(), s.user.LastName.String())

	s.mock.ExpectQuery(sqlSelectWithWhere).WillReturnRows(rows)

	_, err := s.repository.GetUser(context.Background(), s.user.ID.String())

	assert.Nilf(s.T(), err, "%v Should be nil", err)
}

func (s *UserSuite) TestGetUserNotFound() {
	s.mock.ExpectQuery(sqlSelectWithWhere).WillReturnRows(sqlmock.NewRows(nil))

	user, err := s.repository.GetUser(context.Background(), s.user.ID.String())

	assert.NotNil(s.T(), err, "%v Should not be nil", err)
	assert.Nilf(s.T(), user, "%v should be nil", user)
}

func (s *UserSuite) TestDeleteUser() {
	sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow(s.user.ID.String(), s.user.Username.String(), s.user.FirstName.String(), s.user.LastName.String())

	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlDelete).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repository.DeleteUser(context.Background(), s.user.ID.String())

	assert.Nilf(s.T(), err, "%v Should be nil", err)
}

func (s *UserSuite) TestDeleteUserWithError() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlDelete).WillReturnResult(sqlmock.NewErrorResult(errors.New("not found")))
	s.mock.ExpectRollback()

	err := s.repository.DeleteUser(context.Background(), s.user.ID.String())

	assert.NotNilf(s.T(), err, "%v Should not be nil", err)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
