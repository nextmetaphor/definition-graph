package main

import (
	"github.com/nextmetaphor/definition-graph/api"
	"github.com/nextmetaphor/definition-graph/core"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func GetLogger() zerolog.Logger {
	return logger
}

var (
	logger zerolog.Logger
)

func main() {
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	conn, _ := db.OpenDatabase()

	defer db.CloseDatabase(conn)

	core.LoadNodeClassDefinitions([]string{"../definition/nodeClass"}, "yaml", conn)
	core.LoadNodeDefinitionsWithoutEdges([]string{"../definition/node"}, "yaml", conn)
	core.LoadNodeDefinitionsOnlyEdges([]string{"../definition/node"}, "yaml", conn)

	api.Listen(conn)
}
