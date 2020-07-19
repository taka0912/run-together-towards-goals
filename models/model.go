package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
	"os"
)

func Open() *gorm.DB {
	DBMS := "mysql"
	var dataSource string

	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		dataSource = convertDataSource(os.Getenv("CLEARDB_DATABASE_URL"))
	} else {
		dataSource = "root:pass@tcp(mysql:3306)/my_goal"
	}

	db, err := gorm.Open(DBMS, dataSource + "?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func convertDataSource(ds string) (result string) {
	url, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", url.User.String(), url.Host, url.Path)
	return result
}