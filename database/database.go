package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

func IniDatabase() {

	var err error
	DBConn, err = gorm.Open("sqlite3", "./db/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return
	
}