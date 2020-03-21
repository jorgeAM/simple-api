package db

import (
	"log"

	"github.com/jinzhu/gorm"
	// driver to connect with mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetConnection return connection tu database
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Something got wrong to open connection to database")
		return nil
	}

	return db
}
