package mysql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/jorgeAM/api/internal/user/domain"
)

const (
	sqlInsert          = "^INSERT INTO `users` (.+) VALUES"
	sqlSelectAll       = "^SELECT (.+) FROM `users`$"
	sqlSelectWithWhere = "^SELECT (.+) FROM `users` WHERE (.+) ORDER BY `users`.`id` ASC LIMIT 1$"
	sqlUpdate          = "^UPDATE `users` SET (.+) WHERE (.+)"
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

	s.DB.LogMode(true)

	s.repository = NewUserRepository(s.DB)

	s.user = &domain.User{
		ID:        "88109b71-996c-42cd-997e-cbf81cf8f886",
		Username:  "jorgeAM",
		FirstName: "jorge",
		LastName:  "alfaro",
	}
}

func (s *UserSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserSuite) TestNewUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlInsert).WithArgs(
		s.user.ID,
		s.user.Username,
		s.user.FirstName,
		s.user.LastName,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	err := s.repository.NewUser(s.user)

	assert.Nilf(s.T(), err, "%v Should be nil", err)
}

func (s *UserSuite) TestNewUserWithError() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlInsert).WithArgs(
		s.user.ID,
		s.user.Username,
		s.user.FirstName,
		s.user.LastName,
	).WillReturnError(errors.New("Something got wrong to save record"))
	s.mock.ExpectRollback()

	err := s.repository.NewUser(s.user)

	assert.NotNilf(s.T(), err, "%v should not be nil", err)
}

func (s *UserSuite) TestGetUsers() {
	rows := sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow("88109b71-996c-42cd-997e-cbf81cf8f885", "jorgeAM", "jorge", "alfaro").
		AddRow("88109b71-996c-42cd-997e-cbf81cf8f881", "liliMA", "liliana", "murga")

	usersExpected := []*domain.User{
		{
			ID:        "88109b71-996c-42cd-997e-cbf81cf8f885",
			Username:  "jorgeAM",
			FirstName: "jorge",
			LastName:  "alfaro",
		},
		{
			ID:        "88109b71-996c-42cd-997e-cbf81cf8f881",
			Username:  "liliMA",
			FirstName: "liliana",
			LastName:  "murga",
		},
	}

	s.mock.ExpectQuery(sqlSelectAll).WillReturnRows(rows)

	users, err := s.repository.GetUsers()

	assert.Nilf(s.T(), err, "%v Should be nil", err)
	assert.Equal(s.T(), usersExpected, users)
}

func (s *UserSuite) TestGetUsersWithError() {
	s.mock.ExpectQuery(sqlSelectAll).WillReturnError(errors.New("something got wrong"))

	users, err := s.repository.GetUsers()

	assert.NotNilf(s.T(), err, "%v Should not be nil", err)
	assert.Nilf(s.T(), users, "%v Should be nil", users)
}

func (s *UserSuite) TestGetUser() {
	rows := sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow(s.user.ID, s.user.Username, s.user.FirstName, s.user.LastName)

	s.mock.ExpectQuery(sqlSelectWithWhere).WillReturnRows(rows)

	user, err := s.repository.GetUser(s.user.ID)

	assert.Nilf(s.T(), err, "%v Should be nil", err)
	assert.Equal(s.T(), s.user, user)
}

func (s *UserSuite) TestGetUserNotFound() {
	s.mock.ExpectQuery(sqlSelectWithWhere).WillReturnRows(sqlmock.NewRows(nil))

	user, err := s.repository.GetUser(s.user.ID)

	assert.NotNil(s.T(), err, "%v Should not be nil", err)
	assert.Nilf(s.T(), user, "%v should be nil", user)
}

func (s *UserSuite) TestDeleteUser() {
	sqlmock.NewRows([]string{"id", "userName", "firstName", "lastName"}).
		AddRow(s.user.ID, s.user.Username, s.user.FirstName, s.user.LastName)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlDelete).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repository.DeleteUser(s.user.ID)

	assert.Nilf(s.T(), err, "%v Should be nil", err)
}

func (s *UserSuite) TestDeleteUserWithError() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(sqlDelete).WillReturnResult(sqlmock.NewErrorResult(errors.New("not found")))
	s.mock.ExpectRollback()

	err := s.repository.DeleteUser(s.user.ID)

	assert.NotNilf(s.T(), err, "%v Should not be nil", err)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
