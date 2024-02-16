package main

import (
	"github.com/nextmetaphor/definition-graph/core"
	"github.com/nextmetaphor/definition-graph/db"
)

func main() {
	conn, _ := db.OpenDatabase()
	defer db.CloseDatabase(conn)

	core.LoadNodeClassDefinitions([]string{"."}, "yaml", conn)
}
