package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	Driver = "sqlite3"
	DbName  = "my_goal.db"
)

func Open() *gorm.DB {
	db, err := gorm.Open(Driver, DbName)
	if err != nil {
		panic(err)
	}
	return db
}

