package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DBConn is the database connection object
var (
	DBConn *gorm.DB
)

// IniDatabase initializes the database connection
func IniDatabase() {

	var err error

	// Open a database connection and save the reference to `DBConn` object
	DBConn, err = gorm.Open("sqlite3", "./db/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return
	
}