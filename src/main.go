package main

import (
	"github.com/nextmetaphor/definition-graph/api"
	"github.com/nextmetaphor/definition-graph/core"
	"github.com/nextmetaphor/definition-graph/db"
)

func main() {
	conn, _ := db.OpenDatabase()
	defer db.CloseDatabase(conn)

	core.LoadNodeClassDefinitions([]string{"../definition/nodeClass"}, "yaml", conn)
	core.LoadNodeDefinitionsWithoutEdges([]string{"../definition/node"}, "yaml", conn)
	core.LoadNodeDefinitionsOnlyEdges([]string{"../definition/node"}, "yaml", conn)

	api.Listen(conn)
}
