package main

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/src/db"
	"github.com/hariNEzuMI928/run-together-towards-goals/src/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	routes.Handler(db.Init())
}