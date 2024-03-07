package core

import (
	"database/sql"
	"encoding/json"
	db2 "github.com/nextmetaphor/definition-graph/db"
	"github.com/rs/zerolog/log"
)

const (
	logCannotSelectNodeClass      = "cannot select nodeClass"
	logCannotSelectNodeClassGraph = "cannot select nodeClass graph"
	logCannotSelectNodeGraph      = "cannot select node graph"
)

func SelectNodeClasses(db *sql.DB) (b []byte, err error) {
	nodeClasses, err := db2.SelectNodeClass(db)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeClass)
		return
	}

	b, err = json.Marshal(nodeClasses)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeClass)
		return
	}

	return
}

func SelectNodes(db *sql.DB, nodeClassNamespace string, nodeClass string) (b []byte, err error) {
	nodeClasses, err := db2.SelectNodes(db, nodeClass, nodeClassNamespace)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeClass)
		return
	}

	b, err = json.Marshal(nodeClasses)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeClass)
		return
	}

	return
}

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

func SelectNodeGraph(db *sql.DB) (b []byte, err error) {
	graph, err := db2.SelectNodeGraph(db)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeGraph)
		return
	}

	b, err = json.Marshal(graph)
	if err != nil {
		log.Error().Err(err).Msg(logCannotSelectNodeGraph)
		return
	}

	return
}
