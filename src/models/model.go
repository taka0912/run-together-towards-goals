package models

import (
	"github.com/jinzhu/gorm"
	"os"
)

func Open() *gorm.DB {
	DBMS := "mysql"
	var USER string
	var PASS string
	var HOST string
	var DBNAME string

	if os.Getenv("DATABASE_URL") != "" {
		// Heroku用
		USER = "b6ccb4c2b8f823"
		PASS = "378638a9"
		HOST = "us-cdbr-east-02.cleardb.com"
		DBNAME = "heroku_114022f1f0797d9"
	} else {
		// ローカル用
		USER = "root"
		PASS = "pass"
		HOST = "mysql"
		DBNAME = "my_goal"
	}

	dataSource := USER + ":" + PASS + "@tcp(" + HOST + ":3306)/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, dataSource)
	if err != nil {
		panic(err.Error())
	}
	return db
}
