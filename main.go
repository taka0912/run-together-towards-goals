package main

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/db"
	"github.com/hariNEzuMI928/run-together-towards-goals/routes"
)

func main() {
	routes.Handler(db.Init())
}

// TODO ディレクトリ構成をいい感じに整理して、goのソースはdokcer/app配下にしまいたい。
