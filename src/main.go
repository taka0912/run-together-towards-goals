package main

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/db"
	"github.com/hariNEzuMI928/run-together-towards-goals/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	routes.Handler(db.Init())
}