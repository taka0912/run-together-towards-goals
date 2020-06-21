package main

import (
	"github.com/daisuzuki829/run-together-towards-goals/db"
	"github.com/daisuzuki829/run-together-towards-goals/routes"
)

func main() {
	routes.Handler(db.Init())
}
