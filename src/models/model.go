package models

import (
	"github.com/jinzhu/gorm"
)

const (
	Driver = "sqlite3"
	DbName  = "my_goal.db"
)

func Open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "pass"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "my_goal"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}


