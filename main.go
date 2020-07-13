package main

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/db"
	"github.com/hariNEzuMI928/run-together-towards-goals/routes"
)

func main() {
	routes.Handler(db.Init())
}
