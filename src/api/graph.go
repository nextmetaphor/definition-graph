package api

import (
	"github.com/nextmetaphor/definition-graph/db"
	"net/http"
)

const (
	logCannotSelectNodeClassGraph = "cannot select nodeClass graph"
	logCannotSelectNodeGraph      = "cannot select node graph"
)

func nodeClassGraphHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db.SelectNodeClassGraph(dbConn)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeClassGraph)
}

func nodeGraphHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db.SelectNodeGraph(dbConn)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeGraph)
}
