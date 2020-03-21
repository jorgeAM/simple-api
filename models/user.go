package models

// User -> model
type User struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"column:userName"`
	FirstName string `json:"firstname" gorm:"column:firstName"`
	LastName  string `json:"lastname" gorm:"column:lastName"`
}
