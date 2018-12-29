package db

import "github.com/jinzhu/gorm"

var DB *gorm.DB
var err error

func Open() error {
   DB, err = gorm.Open("sqlite3", "test.db")
   return err
}

func Close() error {
  return DB.Close()
}