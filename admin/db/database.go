package db

import "github.com/jinzhu/gorm"

// DB is used for other packages to access the database
var DB *gorm.DB
var err error

// Open is used to open the database connection
func Open() error {
	DB, err = gorm.Open("sqlite3", "test.db")
	return err
}

// Close is used to close the database connection
func Close() error {
	return DB.Close()
}
