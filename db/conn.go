package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	// driver to connect with mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	engine = os.Getenv("DB_ENGINE")
	dsn    = os.Getenv("JAWSDB_URL")
)

// GetConnection return connection tu database
func GetConnection() *gorm.DB {
	db, err := gorm.Open(engine, dsn)

	if err != nil {
		log.Fatal("something got wrong to connect with database", err)
		return nil
	}

	return db
}
