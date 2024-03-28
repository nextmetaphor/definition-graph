package api

import (
	db2 "github.com/nextmetaphor/definition-graph/db"
	"net/http"
)

const (
	logCannotSelectNodeClassGraph = "cannot select nodeClass graph"
	logCannotSelectNodeGraph      = "cannot select node graph"
)

func nodeClassGraphHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db2.SelectNodeClassGraph(db)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeClassGraph)
}

func nodeGraphHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db2.SelectNodeGraph(db)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeGraph)
}
