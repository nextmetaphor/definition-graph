package core

import (
	"database/sql"
	"encoding/json"
	db2 "github.com/nextmetaphor/definition-graph/db"
	"github.com/rs/zerolog/log"
)

const (
	logCannotSelectNodeClassGraph = "cannot select nodeclass graph"
)

func SelectNodeClassGraph(db *sql.DB) (b []byte, err error) {
	graph, err := db2.SelectNodeClassGraph(db)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeClassGraph)
		return
	}

	b, err = json.Marshal(graph)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeClassGraph)
		return
	}

	return
}
