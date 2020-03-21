package db

import (
	"os"

	"github.com/jinzhu/gorm"
	// driver to connect with mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	engine   = os.Getenv("DB_ENGINE")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DATABASE")
)

// GetConnection return connection tu database
func GetConnection() (*gorm.DB, error) {
	dns := user + ":" + password +
		"@(" + host + ":" + port + ")/" +
		database + "?charset=utf8&parseTime=True&loc=Local"
	return gorm.Open(engine, dns)
}
