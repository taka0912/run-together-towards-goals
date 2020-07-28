package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/url"
	"os"
)

func Open() *gorm.DB {
	driver := "mysql"
	var dataSource string

	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		dataSource = convertDataSource(os.Getenv("CLEARDB_DATABASE_URL"))
	} else {
		dataSource = "root:pass@tcp(mysql:3306)/my_goal"
	}
	databaseConnect :=  dataSource+"?parseTime=true&charset=utf8"
	db, err := gorm.Open(driver, databaseConnect)
	if err != nil {
		panic(err.Error())
	}
	_ = os.Setenv("DATABASE_URL", databaseConnect)
	return db
}

func convertDataSource(ds string) (result string) {
	parse, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", parse.User.String(), parse.Host, parse.Path)
	return result
}