package main

import (
	"github.com/daisuzuki829/run_together_towards_goals/db"
	"github.com/daisuzuki829/run_together_towards_goals/routes"
)

func main() {
	routes.Handler(db.Init())
}