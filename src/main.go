package main

import (
	"github.com/nextmetaphor/definition-graph/api"
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

	api.Listen(conn)
}
